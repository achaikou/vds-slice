package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/equinor/oneseismic-api/internal/cache"
	"github.com/equinor/oneseismic-api/internal/core"
)

func httpStatusCode(err error) int {
	switch err.(type) {
	case *core.InvalidArgument:
		return http.StatusBadRequest
	case *core.InternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

/* Call abortOnError on the context in case of an error
 *
 * This function is designed specifically for our endpoint handler functions
 * and aims at making the error handling as short and concise as possible.
 *
 * If err != nil the error will be mapped to an appropriate http status code
 * through the httpStatusCode mapper, and ctx.AbortWithError will be called
 * with this status and the error itself. It then returns true to indicate that
 * the context have been aborted.
 *
 * If err == nil the ctx is left untouched and this function returns false,
 * indicating that the context was not aborted.
 *
 * The result is a one line error handling:
 *
 *     err, _ := func()
 *     if abortOnError(ctx, err) { return }
 */
func abortOnError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	ctx.AbortWithError(httpStatusCode(err), err)

	return true
}

type Endpoint struct {
	MakeVdsConnection core.ConnectionMaker
	Cache             cache.Cache
}

func prepareRequestLogging(ctx *gin.Context, request Stringable) {
	// ignore possible errors as they should not change outcome for the user
	requestString, _ := request.toString()
	ctx.Set("request", requestString)
}

func prepareMetricsLogging(ctx *gin.Context, request interface {
	credentials() ([]string, []string, string)
}) {
	vds, _, _ := request.credentials()
	ctx.Set("vds", vds)
}

func (e *Endpoint) makeDataRequest(
	ctx *gin.Context,
	request DataRequest,
) {
	prepareRequestLogging(ctx, request)
	prepareMetricsLogging(ctx, request)

	connections, binaryOperator, err := e.readConnectionParameters(ctx, request.getRequestedResource())
	if err != nil {
		return
	}

	cacheKey, err := request.hash()
	if abortOnError(ctx, err) {
		return
	}

	cacheEntry, hit := e.Cache.Get(cacheKey)
	if hit {
		var isAuthorizedToRead = true
		for i := 0; i < len(connections); i++ {
			if !connections[i].IsAuthorizedToRead() {
				isAuthorizedToRead = false
			}
		}
		if isAuthorizedToRead {
			ctx.Set("cache-hit", true)
			writeResponse(ctx, cacheEntry.Metadata(), cacheEntry.Data())
			return
		}
	}

	handle, err := core.CreateDSHandle(connections, binaryOperator)
	if abortOnError(ctx, err) {
		return
	}
	defer handle.Close()

	data, metadata, err := request.execute(handle)
	if abortOnError(ctx, err) {
		return
	}

	e.Cache.Set(cacheKey, cache.NewCacheEntry(data, metadata))

	writeResponse(ctx, metadata, data)
}

func (e *Endpoint) readConnectionParameters(
	ctx *gin.Context,
	request RequestedResource,
) ([]core.Connection, uint32, error) {

	vdsUrls, sasTokens, binaryOperatorString := request.credentials()

	binaryOperator, err := core.GetBinaryOperator(binaryOperatorString)
	if abortOnError(ctx, err) {
		return nil, core.BinaryOperatorInvalidOperator, err
	}

	var connections []core.Connection

	if len(vdsUrls) == 1 && binaryOperator != core.BinaryOperatorNoOperator {
		err := core.NewInvalidArgument("Binary operator must be empty when a single VDS url is provided")
		if abortOnError(ctx, err) {
			return nil, core.BinaryOperatorInvalidOperator, err
		}
	} else if len(vdsUrls) == 2 && binaryOperator == core.BinaryOperatorNoOperator {
		err := core.NewInvalidArgument("Binary operator must be provided when two VDS urls are provided")
		if abortOnError(ctx, err) {
			return nil, core.BinaryOperatorInvalidOperator, err
		}
	} else if len(vdsUrls) > 2 {
		err := core.NewInvalidArgument("No endpoint accepts more than two vds urls.")
		if abortOnError(ctx, err) {
			return nil, core.BinaryOperatorInvalidOperator, err
		}
	}

	for i := 0; i < len(vdsUrls); i++ {
		vdsConn, err := e.MakeVdsConnection(vdsUrls[i], sasTokens[i])
		if abortOnError(ctx, err) {
			return nil, core.BinaryOperatorInvalidOperator, err
		}
		connections = append(connections, vdsConn)
	}

	return connections, binaryOperator, nil
}

func parseGetRequest(ctx *gin.Context, v Normalizable) error {
	query, status := ctx.GetQuery("query")
	if !status {
		return core.NewInvalidArgument(
			"GET request to specified endpoint requires a 'query' parameter",
		)
	}
	if err := json.Unmarshal([]byte(query), v); err != nil {
		msg := "Please ensure that the supplied query is valid " +
			"and conforms to the expected swagger Request specification: %v"
		return core.NewInvalidArgument(
			fmt.Sprintf(msg, err.Error()))
	}

	if err := binding.Validator.ValidateStruct(v); err != nil {
		return core.NewInvalidArgument(err.Error())
	}

	return v.NormalizeConnection()
}

func parsePostRequest(ctx *gin.Context, v Normalizable) error {
	if err := ctx.ShouldBind(v); err != nil {
		return core.NewInvalidArgument(err.Error())
	}
	return v.NormalizeConnection()
}
