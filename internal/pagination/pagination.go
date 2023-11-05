package pagination

const LimitPerPage = 10

type PaginatedData[T any] struct {
	CurentPage    int
	AmountOfPages int
	Data          []T
}

type PaginationRequest struct {
	PerPageAmount int
	CurrentPage   int
}

// 10 * (1 - 1)
// 10 * (2 - 1)
// perpage = 10
// page = 1, offset = 0
// page = 2, offset = 10
// page = 3, offset = 20
func CalcOffset(pgReq PaginationRequest) int {
	return (pgReq.PerPageAmount * (pgReq.CurrentPage - 1))

}
