package cocktail

type CocktailDelegate interface {

	/** Let delegate decides whether a request should be handle or not. */
	ShouldServeHTTP(c *Context) bool

	/** Notify delegate that a request had been served. */
	DidServeHTTP(c *Context)
	/** Notify delegate that a request will be served. */
	WillServeHTTP(c *Context)
}
