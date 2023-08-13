# grpc-pagination

## Summary 
We are strongly relying on GRPC as the communication in our services. APIs often need to provide data collection, most 
commonly in the List standard method. It is important that collections be paginated for performance, user experience, 
scalability, and resource management. 

Since this is going to be needed in several services we create this library, so we standardized it and save time 
for further developments. 

### AIP-158

This library implements Google [API Improvement Proposals (AIP) 158](https://google.aip.dev/158). 
As a summary, these are the most important things about it, but is strongly recommend to read the AIP for more detail: 
* Request messages for collections should define an int32 page_size field but must not be required.
* Request messages for collections should define a string page_token field, allowing users to advance to the next page in the collection.
* Response messages for collections should define a string next_page_token field, providing the user with a page token that may be used to retrieve the next page.
* The request definition for a paginated operation may define an int32 skip field to allow the user to skip results.
* Page tokens provided by APIs must be opaque (but URL-safe) strings, and must not be user-parseable.
* The user is expected to keep all arguments to the RPC the same (only can change page_token and  page_size); if any arguments are different, the API should send an INVALID_ARGUMENT error.

### Stable pagination

The library is meant for using [Stable Pagination](https://morningcoffee.io/stable-pagination.html).
As a summary, these are the most important things about it, but is strongly recommend to read the blog for more detail:
* We should store the LastID inside the token, so we can do Cursor Based Pagination:
  * If the id is an auto-incremental we can simply do `WHERE id < LastID` 
  * If the id is a ntwrk guid we can parse the time from it for doing something like `WHERE created_at < LastIDCreatedAt and GUID < LastGuid` 
  * If the id is simply a uuid and doesn’t contain the time inside, this won’t work because we are not storing the created_at time. This may change in the future anyway.

## Installing

```bash
$ go get -u github.com/ntwrk1/grpc-pagination
```

## Usage

### [AIP-132](https://google.aip.dev/132) (Standard method: List)

-	Use [`pagination.PageToken`](./pagetoken.go) to implement offset-based pagination.

  ``` go
  package awesomeProject

import (
	"context"
	"strings"

	"github.com/ntwrk1/grpc-pagination"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Storage Storage
}

func (s *Server) ListShelves(ctx context.Context, request *library.ListShelvesRequest) (*library.ListShelvesResponse, error) {
	// Handle request constraints.
	const (
		maxPageSize     = 1000
		defaultPageSize = 100
	)
	switch {
	case request.PageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size is negative")
	case request.PageSize == 0:
		request.PageSize = defaultPageSize
	case request.PageSize > maxPageSize:
		request.PageSize = maxPageSize
	}
	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}
	// Query the storage.
	result, err := s.Storage.ListShelves(ctx, &ListShelvesQuery{
		Cursor:   pageToken.LastID,
		PageSize: request.GetPageSize(),
	})
	if err != nil {
		return nil, err
	}
	// Build the response.
	response := &library.ListShelvesResponse{
		Shelves: result.Shelves,
	}
	// Set the next page token.
	if result.HasNextPage && len(result.Shelves) > 0 {
		lastName := result.Shelves[len(result.Shelves)-1].Name
		// Shelf names have the form `shelves/{shelf_id}`.
		lastID := strings.Replace(lastName, "shelves/", "", 0)
		response.NextPageToken = pageToken.Next(lastID).String()
	}
	// Respond.
	return response, nil
}

type Storage struct {
}

type ListShelvesQuery struct {
	Cursor   string
	PageSize int32
}

type ListShelvesResponse struct {
	Shelves     []*library.Shelf
	HasNextPage bool
}

func (s *Storage) ListShelves(ctx context.Context, request *ListShelvesQuery) (ListShelvesResponse, error) {
	shelves := make([]*library.Shelf, 0)

	if request.Cursor != "" {
        // If cursor is a NTWRK GUID
        id, err := guid.ParseString()
        if err != nil {
            return errors.New("parsing cursor string into a guid")
        }
        // Order by CreateTime ASC or DESC.
        // Do a predicate like shelf.CreateTimeLT(id.Time()) if order DESC or CreateTimeGT order ASC.

        // If cursor is an auto incremental ID
        idNumeric, err := strconv.Atoi(request.Cursor)
        if err != nil {
            return errors.New("parsing cursor string into an integer")
        }
        // Order by Id ASC or DESC.
        // Do a predicate like shelf.IDLT(idNumeric) if order DESC or IDGT order ASC.  
	}

	// Remember to limit the results based on request.PageSize

	return ListShelvesResponse{
		HasNextPage: pagination.HasNextPage(request.PageSize, len(shelves)),
	}, nil
}
  ```


### Changelog

#### 1.0.0 (May 24, 2023)
- Initial Release

### Notes
This is a kind of fork of https://github.com/einride/aip-go.
