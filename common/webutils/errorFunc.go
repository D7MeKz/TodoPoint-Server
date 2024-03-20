package webutils

//type errorResponse struct {
//	Code    int    `json:"code"`
//	Error   string `json:"error"`
//	Message string `json:"message"`
//}
//
//func InvalidDataError(ctx *gin.Context, err error) {
//	res := errorResponse{Code: http.StatusUnauthorized, Error: err.Error()}
//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
//}
//
//func InternalDBError(ctx *gin.Context, err error) {
//	res := errorResponse{Code: http.StatusInternalServerError, Error: err.Error()}
//	ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
//}
//
//func BadRequestError(ctx *gin.Context, err error) {
//	res := errorResponse{Code: http.StatusBadRequest, Error: err.Error()}
//	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
//}
