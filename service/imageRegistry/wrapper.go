package imageRegistry

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	. "github.com/volcengine/volc-sdk-golang/service/imageRegistry/models"
)

func (p *ImageRegistry) CreateNamespaceBasic(req *CreateNamespaceBasicRequest) (*CreateNamespaceBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateNamespaceBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateNamespaceBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("CreateNamespaceBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateNamespaceBasic: fail to do request, %v", err)
			}
			result := new(CreateNamespaceBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateNamespaceBasic: fail to do request, %v", err)
	}
	result := new(CreateNamespaceBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) ValidateNamespaceBasic(req *ValidateNamespaceBasicRequest) (*ValidateNamespaceBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ValidateNamespaceBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ValidateNamespaceBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ValidateNamespaceBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ValidateNamespaceBasic: fail to do request, %v", err)
			}
			result := new(ValidateNamespaceBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ValidateNamespaceBasic: fail to do request, %v", err)
	}
	result := new(ValidateNamespaceBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) DeleteNamespaceBasic(req *DeleteNamespaceBasicRequest) (*DeleteNamespaceBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteNamespaceBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteNamespaceBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("DeleteNamespaceBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteNamespaceBasic: fail to do request, %v", err)
			}
			result := new(DeleteNamespaceBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteNamespaceBasic: fail to do request, %v", err)
	}
	result := new(DeleteNamespaceBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) GetNamespaceBasic(req *GetNamespaceBasicRequest) (*GetNamespaceBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNamespaceBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetNamespaceBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetNamespaceBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetNamespaceBasic: fail to do request, %v", err)
			}
			result := new(GetNamespaceBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetNamespaceBasic: fail to do request, %v", err)
	}
	result := new(GetNamespaceBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) ListNamespacesBasic(req *ListNamespacesBasicRequest) (*ListNamespacesBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListNamespacesBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ListNamespacesBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ListNamespacesBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ListNamespacesBasic: fail to do request, %v", err)
			}
			result := new(ListNamespacesBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ListNamespacesBasic: fail to do request, %v", err)
	}
	result := new(ListNamespacesBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) CreateRepositoryBasic(req *CreateRepositoryBasicRequest) (*CreateRepositoryBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateRepositoryBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateRepositoryBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("CreateRepositoryBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateRepositoryBasic: fail to do request, %v", err)
			}
			result := new(CreateRepositoryBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateRepositoryBasic: fail to do request, %v", err)
	}
	result := new(CreateRepositoryBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) ValidateRepositoryBasic(req *ValidateRepositoryBasicRequest) (*ValidateRepositoryBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ValidateRepositoryBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ValidateRepositoryBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ValidateRepositoryBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ValidateRepositoryBasic: fail to do request, %v", err)
			}
			result := new(ValidateRepositoryBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ValidateRepositoryBasic: fail to do request, %v", err)
	}
	result := new(ValidateRepositoryBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) DeleteRepositoryBasic(req *DeleteRepositoryBasicRequest) (*DeleteRepositoryBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteRepositoryBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteRepositoryBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("DeleteRepositoryBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteRepositoryBasic: fail to do request, %v", err)
			}
			result := new(DeleteRepositoryBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteRepositoryBasic: fail to do request, %v", err)
	}
	result := new(DeleteRepositoryBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) UpdateRepositoryBasic(req *UpdateRepositoryBasicRequest) (*UpdateRepositoryBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateRepositoryBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdateRepositoryBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("UpdateRepositoryBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdateRepositoryBasic: fail to do request, %v", err)
			}
			result := new(UpdateRepositoryBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdateRepositoryBasic: fail to do request, %v", err)
	}
	result := new(UpdateRepositoryBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) GetRepositoryBasic(req *GetRepositoryBasicRequest) (*GetRepositoryBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetRepositoryBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetRepositoryBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetRepositoryBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetRepositoryBasic: fail to do request, %v", err)
			}
			result := new(GetRepositoryBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetRepositoryBasic: fail to do request, %v", err)
	}
	result := new(GetRepositoryBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) ListRepositoriesBasic(req *ListRepositoriesBasicRequest) (*ListRepositoriesBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListRepositoriesBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ListRepositoriesBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ListRepositoriesBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ListRepositoriesBasic: fail to do request, %v", err)
			}
			result := new(ListRepositoriesBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ListRepositoriesBasic: fail to do request, %v", err)
	}
	result := new(ListRepositoriesBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) DeleteTagBasic(req *DeleteTagBasicRequest) (*DeleteTagBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteTagBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteTagBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("DeleteTagBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteTagBasic: fail to do request, %v", err)
			}
			result := new(DeleteTagBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteTagBasic: fail to do request, %v", err)
	}
	result := new(DeleteTagBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) GetTagBasic(req *GetTagBasicRequest) (*GetTagBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTagBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetTagBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetTagBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetTagBasic: fail to do request, %v", err)
			}
			result := new(GetTagBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetTagBasic: fail to do request, %v", err)
	}
	result := new(GetTagBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) GetTagAdditionBasic(req *GetTagAdditionBasicRequest) (*GetTagAdditionBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTagAdditionBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetTagAdditionBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetTagAdditionBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetTagAdditionBasic: fail to do request, %v", err)
			}
			result := new(GetTagAdditionBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetTagAdditionBasic: fail to do request, %v", err)
	}
	result := new(GetTagAdditionBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) ListTagsBasic(req *ListTagsBasicRequest) (*ListTagsBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListTagsBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ListTagsBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ListTagsBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ListTagsBasic: fail to do request, %v", err)
			}
			result := new(ListTagsBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ListTagsBasic: fail to do request, %v", err)
	}
	result := new(ListTagsBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ImageRegistry) GetAuthorizationTokenBasic(req *GetAuthorizationTokenBasicRequest) (*GetAuthorizationTokenBasicResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetAuthorizationTokenBasicRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetAuthorizationTokenBasic", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetAuthorizationTokenBasic", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetAuthorizationTokenBasic: fail to do request, %v", err)
			}
			result := new(GetAuthorizationTokenBasicResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAuthorizationTokenBasic: fail to do request, %v", err)
	}
	result := new(GetAuthorizationTokenBasicResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
