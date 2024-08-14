import torch
from cuda import cuda

if __name__ == "__main__":
    err = cuda.cuInit(0)[0]
    assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)
    a = torch.rand((2, 3), device='cuda')
    err, sm_count = cuda.cuCtxGetExecAffinity(
        cuda.CUexecAffinityType.CU_EXEC_AFFINITY_TYPE_SM_COUNT
    )
    assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)
    print(sm_count)

    sm_count.param.smCount.val = 100
    err, comp_ctx = cuda.cuCtxCreate_v3([sm_count], 1, 0, 0)
    sm_count.param.smCount.val = 25
    err, mem_ctx = cuda.cuCtxCreate_v3([sm_count], 1, 0, 0)
    assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)

    s1 = torch.cuda.Stream()
    s2 = torch.cuda.Stream()
    with torch.cuda.stream(s1):
        err = cuda.cuCtxSetCurrent(comp_ctx)[0]
        err, sm_count = cuda.cuCtxGetExecAffinity(
            cuda.CUexecAffinityType.CU_EXEC_AFFINITY_TYPE_SM_COUNT
        )
        assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)
        print(sm_count)

    with torch.cuda.stream(s2):
        err = cuda.cuCtxSetCurrent(mem_ctx)[0]
        err, sm_count = cuda.cuCtxGetExecAffinity(
            cuda.CUexecAffinityType.CU_EXEC_AFFINITY_TYPE_SM_COUNT
        )
        assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)
        print(sm_count)

    with torch.cuda.stream(s1):
        err, sm_count = cuda.cuCtxGetExecAffinity(
            cuda.CUexecAffinityType.CU_EXEC_AFFINITY_TYPE_SM_COUNT
        )
        assert err == cuda.CUresult.CUDA_SUCCESS, cuda.cuGetErrorString(err)
        print(sm_count)

    torch.cuda.synchronize()
