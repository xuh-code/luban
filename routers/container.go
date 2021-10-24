package routers

import (
	"github.com/gin-gonic/gin"
	"pigs/controller/k8s/cluster"
	"pigs/controller/k8s/evnet"
	"pigs/controller/k8s/node"
)

func InitContainerRouter(r *gin.RouterGroup) {
	K8sClusterRouter := r.Group("k8s")
	{
		K8sClusterRouter.POST("cluster", cluster.CreateK8SCluster)
		K8sClusterRouter.GET("cluster", cluster.ListK8SCluster)
		K8sClusterRouter.GET("cluster/secret", cluster.ClusterSecret)
		K8sClusterRouter.POST("delCluster", cluster.DelK8SCluster)
		K8sClusterRouter.GET("cluster/detail", cluster.GetK8SClusterDetail)
		K8sClusterRouter.GET("events", evnet.Events)
		K8sClusterRouter.GET("node", node.GetNodes)
		K8sClusterRouter.GET("node/detail", node.GetNodeDetail)
		K8sClusterRouter.POST("node/schedule", node.NodeUnschedulable)
	}
}
