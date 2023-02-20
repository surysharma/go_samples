An example of strategy pattern to showcase the importance of Open closed principal.

In the example, we have a number of products, with attributes
like name, description and color.

We want to filter out the products based on the
color or name.

Instead of adding filtering logic with 
if conditions, A `FilterStrategy` is used to delegate
the business logic in separate functions, each 
implementing the FilterStrategy.isPresent contract.

Usage

`go run main.go`