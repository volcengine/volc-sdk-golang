package imageRegistry

import (
	"fmt"
	"testing"

	. "github.com/volcengine/volc-sdk-golang/service/imageRegistry/models"
)

const (
	Ak = "" // write your access key
	Sk = "" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func ListNamespacesBasic() {
	res, err := DefaultInstance.ListNamespacesBasic(&ListNamespacesBasicRequest{})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func ListRepositoriesBasic(namespace string) {
	res, err := DefaultInstance.ListRepositoriesBasic(&ListRepositoriesBasicRequest{})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func ListTagsBasic(namespace string, repo string, tagType string) {
	res, err := DefaultInstance.ListTagsBasic(&ListTagsBasicRequest{
		Namespace: namespace,
		Repo:      repo,
		Type:      tagType,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetTagBasic(namespace string, repo string, name string) {
	res, err := DefaultInstance.GetTagBasic(&GetTagBasicRequest{
		Namespace: namespace,
		Repo:      repo,
		Name:      name,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetTagAdditionBasic(namespace string, repo string, name string, digest string, addition string) {
	res, err := DefaultInstance.GetTagAdditionBasic(&GetTagAdditionBasicRequest{
		Namespace: namespace,
		Repo:      repo,
		Name:      name,
		Digest:    digest,
		Addition:  addition,
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestImageRegistry_ListNamespacesBasic(t *testing.T) {
	ListNamespacesBasic()
}

func TestImageRegistry_ListRepositoriesBasic(t *testing.T) {
	ListRepositoriesBasic("xfl-test")
}

func TestImageRegistry_ListTagsBasic(t *testing.T) {
	ListTagsBasic("xfl-test", "xfl-test/test-ruby-app", "Image")
}

func TestImageRegistry_GetTagBasic(t *testing.T) {
	GetTagBasic("xfl-test", "xfl-test/test-ruby-app", "latest")
}

func TestImageRegistry_GetTagAdditionBasic(t *testing.T) {
	GetTagAdditionBasic("xfl-test", "xfl-test/test-ruby-app", "latest", "sha256:1e994581cd4a41d04843dd9b3ce6b182eae2794cd6dfd2c6866de59efc1dbbee", "Values")
}
