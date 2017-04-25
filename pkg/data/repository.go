package data

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"fmt"
)

//// Object which holds the values for different metrics for all the resources of an entity
type MetricMap map[ResourceType]map[MetricPropType]*Metric

// Interface for Repository Entity
type RepositoryEntity interface {
	GetId() string
	GetType() proto.EntityDTO_EntityType
	GetResourceMetrics() MetricMap
	GetResourceMetric(resourceType ResourceType, metricType MetricPropType) (*Metric, error)
}

// Interface for a Repository
type Repository interface {
	GetEntity(entityType proto.EntityDTO_EntityType, id string) RepositoryEntity
	GetEntityInstances(entityType proto.EntityDTO_EntityType) []RepositoryEntity
}

type Metric struct {
	value *float64
}
// =============================================== Entity Metrics ======================================

func (resourceMetrics MetricMap) SetResourceMetric(resourceType ResourceType, metricType MetricPropType, value *float64) {
	//resourceMap, exists := resourceMetrics.metricMap[resourceType]
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		resourceMap = make(map[MetricPropType]*Metric)
		resourceMetrics[resourceType] = resourceMap
		//resourceMetrics.metricMap[resourceType] = resourceMap
	}
	metric, ok := resourceMap[metricType]
	if !ok {
		metric = &Metric{}
	}
	metric.value = value
	resourceMap[metricType] = metric
}

func (resourceMetrics MetricMap) GetResourceMetric(resourceType ResourceType, metricType MetricPropType) (*Metric, error) {
	//resourceMap, exists := resourceMetrics.metricMap[resourceType]
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for resource %s\n", resourceType)
		return nil, fmt.Errorf("missing metrics for resource %s", resourceType)
	}
	metric, exists := resourceMap[metricType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", metricType)
		return nil, fmt.Errorf("missing metrics for type %s:%s", resourceType, metricType)
	}
	return metric, nil
}

func (resourceMetrics MetricMap) printMetrics() {
	//glog.Infof("Entity %s\n", resourceMetrics.entityId)
	//fmt.Printf("Entity %s\n", resourceMetrics.entityId)
	for rt, resourceMap := range resourceMetrics { //.metricMap {
		//fmt.Printf("Resource Type %s\n", rt)
		for mkey, metric := range resourceMap {
			if (metric != nil) {
				glog.Infof("\t\t%s::%s : %f\n", rt, mkey, *metric.value)
				fmt.Printf("\t\t%s::%s : %f\n", rt, mkey, *metric.value)
			} else {
				glog.Infof("\t\t%s::%s : %f\n", rt, mkey, metric.value)
				fmt.Printf("\t\t%s::%s : %f\n", rt, mkey, metric.value)

			}
		}
	}
}

func PrintEntity(entity RepositoryEntity) {
	glog.Infof("Entity %s::%s\n", entity.GetType(), entity.GetId())
	fmt.Printf("Entity %s::%s\n", entity.GetType(), entity.GetId())
	resourceMetrics := entity.GetResourceMetrics()
	resourceMetrics.printMetrics()
}

//func PrintRepository(repository Repository) {
//	PrintEntity(repository.GetAgentEntity())
//	taskEntities := repository.GetTaskEntities()
//	for _, taskEntity := range taskEntities {
//		PrintEntity(taskEntity)
//	}
//	containerEntities := repository.GetContainerEntities()
//	for _, containerEntity := range containerEntities {
//		PrintEntity(containerEntity)
//	}
//}
