package registration

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/pkg/supplychain"
)

var (
	respTimeType    = proto.CommodityDTO_RESPONSE_TIME
	transactionType = proto.CommodityDTO_TRANSACTION
	key             = "key-placeholder"

	respTimeTemplateComm *proto.TemplateCommodity = &proto.TemplateCommodity{
		CommodityType: &respTimeType,
		Key:           &key,
	}

	transactionTemplateComm *proto.TemplateCommodity = &proto.TemplateCommodity{
		CommodityType: &transactionType,
		Key:           &key,
	}
)

type SupplyChainFactory struct{}

func (f *SupplyChainFactory) CreateSupplyChain() ([]*proto.TemplateDTO, error) {
	appNode, err := f.buildAppSupplyBuilder()
	if err != nil {
		return nil, err
	}

	vAppNode, err := f.buildVAppSupplyBuilder()
	if err != nil {
		return nil, err
	}

	return supplychain.NewSupplyChainBuilder().Top(vAppNode).Entity(appNode).
		Create()
}

func (f *SupplyChainFactory) buildAppSupplyBuilder() (*proto.TemplateDTO, error) {
	builder := supplychain.NewSupplyChainNodeBuilder(proto.EntityDTO_APPLICATION).
		Sells(transactionTemplateComm).
		Sells(respTimeTemplateComm)
	builder.SetPriority(-1)
	builder.SetTemplateType(proto.TemplateDTO_BASE)
	//builder.SetTemplateType(proto.TemplateDTO_EXTENSION)

	return builder.Create()
}

func (f *SupplyChainFactory) buildVAppSupplyBuilder() (*proto.TemplateDTO, error) {
	builder := supplychain.NewSupplyChainNodeBuilder(proto.EntityDTO_VIRTUAL_APPLICATION).
		Provider(proto.EntityDTO_APPLICATION, proto.Provider_LAYERED_OVER).
		Buys(transactionTemplateComm).
		Buys(respTimeTemplateComm)
	builder.SetPriority(-1)
	builder.SetTemplateType(proto.TemplateDTO_BASE)
	//builder.SetTemplateType(proto.TemplateDTO_EXTENSION)

	return builder.Create()
}
