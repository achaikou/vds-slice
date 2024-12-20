#ifndef ONESEISMIC_API_DATAHANDLE_HPP
#define ONESEISMIC_API_DATAHANDLE_HPP

#include <memory>
#include <string>

#include <OpenVDS/OpenVDS.h>
#include <functional>

#include "metadatahandle.hpp"
#include "subcube.hpp"

using voxel = float[OpenVDS::Dimensionality_Max];

class DataHandle {

public:
    virtual ~DataHandle() {};
    virtual void close() = 0;

    virtual MetadataHandle const& get_metadata() const noexcept(true) = 0;

    virtual std::int64_t samples_buffer_size(std::size_t const nsamples) noexcept(false) = 0;

    virtual void read_samples(
        void* const buffer,
        std::int64_t const size,
        voxel const* samples,
        std::size_t const nsamples,
        enum interpolation_method const interpolation_method
    ) noexcept(false) = 0;

    virtual std::int64_t subcube_buffer_size(SubCube const& subcube) noexcept(false) = 0;

    virtual void read_subcube(
        void* const buffer,
        std::int64_t size,
        SubCube const& subcube
    ) noexcept(false) = 0;

    virtual std::int64_t traces_buffer_size(std::size_t const ntraces) noexcept(false) = 0;

    virtual void read_traces(
        void* const buffer,
        std::int64_t const size,
        voxel const* coordinates,
        std::size_t const ntraces,
        enum interpolation_method const interpolation_method
    ) noexcept(false) = 0;

    static OpenVDS::VolumeDataFormat format() noexcept(true);
};

class SingleDataHandle : public DataHandle {
    SingleDataHandle(OpenVDS::VDSHandle handle);
    friend SingleDataHandle make_single_datahandle(const char* url, const char* credentials);

public:
    void close();

    SingleMetadataHandle const& get_metadata() const noexcept (true);

    static OpenVDS::VolumeDataFormat format() noexcept (true);

    std::int64_t subcube_buffer_size(SubCube const& subcube) noexcept (false);

    void read_subcube(
        void * const buffer,
        std::int64_t size,
        SubCube const& subcube
    ) noexcept (false);

    std::int64_t traces_buffer_size(std::size_t const ntraces) noexcept (false);

    void read_traces(
        void * const                    buffer,
        std::int64_t const              size,
        voxel const*                    coordinates,
        std::size_t const               ntraces,
        enum interpolation_method const interpolation_method
    ) noexcept (false);


    std::int64_t samples_buffer_size(std::size_t const nsamples) noexcept (false);

    void read_samples(
        void * const                    buffer,
        std::int64_t const              size,
        voxel const*                    samples,
        std::size_t const               nsamples,
        enum interpolation_method const interpolation_method
    ) noexcept (false);

private:
    OpenVDS::VDSHandle m_handle;
    OpenVDS::VolumeDataAccessManager m_access_manager;
    SingleMetadataHandle m_metadata;

    static int constexpr lod_level = 0;
    static int constexpr channel = 0;
};

SingleDataHandle make_single_datahandle(
    const char* url,
    const char* credentials
) noexcept(false);

class DoubleDataHandle : public DataHandle {

public:
    DoubleDataHandle(SingleDataHandle handle_a, SingleDataHandle handle_b, enum binary_operator binary_symbol);
    friend DoubleDataHandle make_double_datahandle(
        const char* url_a,
        const char* credentials_a,
        const char* url_b,
        const char* credentials_b,
        enum binary_operator binary_symbol
    );

    void close();

    DoubleMetadataHandle const& get_metadata() const noexcept(true);

    static OpenVDS::VolumeDataFormat format() noexcept(true);

    std::int64_t subcube_buffer_size(SubCube const& subcube) noexcept(false);

    void read_subcube(
        void* const buffer,
        std::int64_t size,
        SubCube const& subcube
    ) noexcept(false);

    std::int64_t traces_buffer_size(std::size_t const ntraces) noexcept(false);

    void read_traces(
        void* const buffer,
        std::int64_t const size,
        voxel const* coordinates,
        std::size_t const ntraces,
        enum interpolation_method const interpolation_method
    ) noexcept(false);

    std::int64_t samples_buffer_size(std::size_t const nsamples) noexcept(false);

    void read_samples(
        void* const buffer,
        std::int64_t const size,
        voxel const* samples,
        std::size_t const nsamples,
        enum interpolation_method const interpolation_method
    ) noexcept(false);

private:
    SingleDataHandle m_datahandle_a;
    SingleDataHandle m_datahandle_b;
    DoubleMetadataHandle m_metadata;
    std::function<void(float*, const float*, std::size_t)> m_binary_operator;

    static int constexpr lod_level = 0;
    static int constexpr channel = 0;

    SubCube offset_bounds(const SubCube subcube, SingleMetadataHandle metadata);
    void extract_continuous_part_of_trace(
        std::vector<float>* source_traces,
        int source_trace_length,
        long start_extract_index,
        int nsamples_to_extract,
        float* target_buffer
    );
};

DoubleDataHandle make_double_datahandle(
    const char* url_a, const char* credentials_a,
    const char* url_b, const char* credentials_b,
    enum binary_operator bin_operator
) noexcept(false);

void inplace_subtraction(
    float* buffer_A,
    const float* buffer_B,
    std::size_t nsamples
) noexcept(true);

void inplace_addition(
    float* buffer_A,
    const float* buffer_B,
    std::size_t nsamples
) noexcept(true);

void inplace_multiplication(
    float* buffer_A,
    const float* buffer_B,
    std::size_t nsamples
) noexcept(true);

void inplace_division(
    float* buffer_A,
    const float* buffer_B,
    std::size_t nsamples
) noexcept(true);

#endif /* ONESEISMIC_API_DATAHANDLE_HPP */
