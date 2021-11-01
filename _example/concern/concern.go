package example_concern

import (
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Sora233/DDBOT/lsp/concern"
	"github.com/Sora233/DDBOT/lsp/concern_type"
	"github.com/Sora233/DDBOT/lsp/mmsg"
	"github.com/Sora233/DDBOT/lsp/registry"
)

var logger = utils.GetModuleLogger("example-concern")

const (
	Site    string            = "example"
	Example concern_type.Type = "example"
)

type exampleStateManager struct {
	*concern.StateManager
}

func (c *exampleStateManager) GetGroupConcernConfig(groupCode int64, id interface{}) concern.IConfig {
	return NewGroupConcernConfig(c.StateManager.GetGroupConcernConfig(groupCode, id))
}

func newExampleStateManager() *exampleStateManager {
	return &exampleStateManager{concern.NewStateManagerWithStringID("example", true)}
}

type exampleConcern struct {
	*exampleStateManager
	notifyChan chan<- concern.Notify
}

func (c *exampleConcern) Site() string {
	return Site
}

func (c *exampleConcern) Start() error {
	err := c.GetStateManager().Start()
	if err != nil {
		return err
	}
	go c.EmitFreshCore(c.Site(), func(ctype concern_type.Type, id interface{}) error {
		groups, _, _, err := c.GetStateManager().
			ListConcernState(func(groupCode int64, _id interface{}, p concern_type.Type) bool {
				return _id == id && p.ContainAny(ctype)
			})
		if err != nil {
			return err
		}
		for _, group := range groups {
			c.notifyChan <- &notify{
				groupCode: group,
				id:        id.(string),
			}
		}
		return nil
	})
	return nil
}

func (c *exampleConcern) Stop() {
	c.GetStateManager().Stop()
}

func (c *exampleConcern) ParseId(s string) (interface{}, error) {
	return s, nil
}

func (c *exampleConcern) Add(ctx mmsg.IMsgCtx, groupCode int64, id interface{}, ctype concern_type.Type) (concern.IdentityInfo, error) {
	_, err := c.GetStateManager().AddGroupConcern(groupCode, id.(string), ctype)
	if err != nil {
		return nil, err
	}
	return c.Get(id)
}

func (c *exampleConcern) Remove(ctx mmsg.IMsgCtx, groupCode int64, id interface{}, ctype concern_type.Type) (concern.IdentityInfo, error) {
	_, err := c.GetStateManager().RemoveGroupConcern(groupCode, id.(string), ctype)
	if err != nil {
		return nil, err
	}
	return c.Get(id)
}

func (c *exampleConcern) List(groupCode int64, ctype concern_type.Type) ([]concern.IdentityInfo, []concern_type.Type, error) {
	_, ids, ctypes, err := c.GetStateManager().ListConcernState(func(_groupCode int64, id interface{}, p concern_type.Type) bool {
		return groupCode == _groupCode && p.ContainAny(ctype)
	})
	if err != nil {
		return nil, nil, err
	}
	ids, ctypes, err = c.GetStateManager().GroupTypeById(ids, ctypes)
	if err != nil {
		return nil, nil, err
	}
	var result []concern.IdentityInfo
	var resultType []concern_type.Type
	for index, id := range ids {
		info, err := c.Get(id)
		if err != nil {
			continue
		}
		result = append(result, info)
		resultType = append(resultType, ctypes[index])
	}
	return result, resultType, nil
}

func (c *exampleConcern) Get(id interface{}) (concern.IdentityInfo, error) {
	return concern.NewIdentity(id, id.(string)), nil
}

func (c *exampleConcern) GetStateManager() concern.IStateManager {
	return c.StateManager
}

func NewConcern() *exampleConcern {
	return &exampleConcern{
		exampleStateManager: newExampleStateManager(),
		notifyChan:          registry.GetNotifyChan(),
	}
}

func init() {
	registry.RegisterConcernManager(NewConcern(), []concern_type.Type{Example})
}
