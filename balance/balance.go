package balance

type Balancer interface {
	Join() error
	Leave() error
}
