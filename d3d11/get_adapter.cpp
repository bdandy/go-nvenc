#include "get_adapter.h"

IDXGIAdapter* GetNvidiaAdapter() {
    IDXGIAdapter* pAdapter;
    IDXGIFactory* pFactory;

    if(FAILED(CreateDXGIFactory(__uuidof(IDXGIFactory) ,(void**)&pFactory)))
    {
        return NULL;
    }

	for (UINT i = 0; pFactory->EnumAdapters(i, &pAdapter) != DXGI_ERROR_NOT_FOUND; ++i) {
		DXGI_ADAPTER_DESC pDesc;
		pAdapter->GetDesc(&pDesc);
		if(pDesc.VendorId == 0x10DE) {
		    pFactory->Release();
			return pAdapter;
		}
		pAdapter->Release();
	}
	pFactory->Release();
	return NULL;
}