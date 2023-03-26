package helpers

import (
	"alta-cookit-be/utils/consts"
	"errors"
	"mime/multipart"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LimitOffsetConvert(page, limit int) (int, int) {
	offset := -1
	if limit > 0 {
		offset = (page - 1) * limit
	}
	return limit, offset
}

func ExtractIDParam(c echo.Context, pathName string) (uint, error) {
	idStr := c.Param(pathName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New(consts.ECHO_InvaildIdParam)
	}
	return uint(id), nil
}

func ExtractPageLimit(c echo.Context) (page int, limit int) {
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1
	}
	limitStr := c.QueryParam("limit")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return -1, -1
	}
	return page, limit
}

func ExtractQueryParams(queryParams url.Values) map[string]interface{} {
	extractedQueryParams := make(map[string]interface{})
	for key, val := range queryParams {
		if key != "page" && key != "limit"{
			extractedQueryParams[key]=val[0]
		}
	}
	return extractedQueryParams
}

func ExtractFile(c echo.Context, key string) (multipart.File, string, error) {
	f, err := c.FormFile(key)
	if err != nil {
		return nil, "", err
	}

	blobFile, err := f.Open()
	if err != nil {
		return nil, "", err
	}
	defer blobFile.Close()

	return blobFile, f.Filename, nil
}

func ExtractMultipleFiles(c echo.Context, key string) ([]multipart.File, []string, error) {
	blobFiles, fileNames := []multipart.File{}, []string{}

	form, err := c.MultipartForm()
	if err != nil {
		return nil, nil, err
	}
	files := form.File[key]

	for _, file := range files {
		blobFile, err := file.Open()
		if err != nil {
			return nil, nil, err
		}
		defer blobFile.Close()

		blobFiles = append(blobFiles, blobFile)
		fileNames = append(fileNames, file.Filename)
	}

	return blobFiles, fileNames, nil
}
