package handler

import (
	"alta-cookit-be/features/followers"
	"alta-cookit-be/middlewares"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type followHander struct {
	srv followers.FollowService
}

func New(srv followers.FollowService) followers.FollowHandler {
	return &followHander{
		srv: srv,
	}
}

// Follow implements followers.FollowHandler
func (fh *followHander) Follow() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID := c.Param("id")
		followingID, _ := strconv.Atoi(uID)
		id, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error from server"})
		}
		if role == "Admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "cannot access credentials data"})
		}

		if id == uint(followingID) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you cannot follow your self"})
		}
		err = fh.srv.Follow(id, uint(followingID))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you already follow this user"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success follow this user",
		})
	}
}

// ShowAllFollower implements followers.FollowHandler
func (fh *followHander) ShowAllFollower() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		dataCore, err := fh.srv.ShowAllFollower(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		result := []ListFollowerResponse{}
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		for _, val := range dataCore {
			result = append(result, ToListFollowerResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all follower users",
		})
	}
}

// ShowAllFollowing implements followers.FollowHandler
func (fh *followHander) ShowAllFollowing() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		dataCore, err := fh.srv.ShowAllFollowing(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		result := []ListFollowingResponse{}
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		for _, val := range dataCore {
			result = append(result, ToListFollowingResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all following users",
		})
	}
}

// Unfollow implements followers.FollowHandler
func (fh *followHander) Unfollow() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID := c.Param("id")
		followingID, _ := strconv.Atoi(uID)
		id, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "error from server"})
		}
		if role == "Admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "cannot access credentials data"})
		}
		if id == uint(followingID) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you cannot unfollow your self"})
		}

		err = fh.srv.Unfollow(id, uint(followingID))
		if err != nil {
			if strings.Contains(err.Error(), "invalid") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "invalid user id, data not found"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success unfollow this user",
		})
	}
}
