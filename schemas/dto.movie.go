package schemas

type CreateMovieRequest struct {
	Title       string  `json:"title"  binding:"required" `
	Description string  `json:"description"  binding:"required" `
	Rating      float64 `json:"rating"  binding:"required" `
	Image       string  `json:"image"  binding:"required" `
}
type CreateMovieRequestTest struct {
	Title       float64 `json:"title"  binding:"required" `
	Description string  `json:"description"  binding:"required" `
	Rating      float64 `json:"rating"  binding:"required" `
	Image       string  `json:"image"  binding:"required" `
}

type UpdateMovieRequest struct {
	Title       string  `json:"title"  binding:"required" `
	Description string  `json:"description"  binding:"required" `
	Rating      float64 `json:"rating"  binding:"required" `
	Image       string  `json:"image"  binding:"required" `
}
