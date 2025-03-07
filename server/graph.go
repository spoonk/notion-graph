package main

type NodeId string

type Node struct {
	Name     string   `json:"name"`
	NodeId   string   `json:"nodeId"`
	Outgoing []NodeId `json:"outgoing"`
}

type Graph struct {
	G map[NodeId]Node `json:"graph"`
}

func (g Graph) addNode(note NotionNote) {
	node := Node{
		Name:   note.Title,
		NodeId: note.ID,
	}

	_, exists := g.G[NodeId(note.ID)]

	if exists {
		return
	}

	ids := []NodeId{}
	for _, id := range note.RelatedIds {
		ids = append(ids, NodeId(id))
	}

	g.G[NodeId(note.ID)] = Node{
		Name:     node.Name,
		NodeId:   note.ID,
		Outgoing: ids,
	}

}
