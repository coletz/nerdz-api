package api

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"strconv"
)

/*
UserPosts handles the request and returns all the posts written
by the specified user
*/
func UserPosts(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var options *nerdz.PostlistOptions
	if options, e = NewPostlistOptions(c.Request()); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: e.Error(),
			Message:      "NewPostlistOptions error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	posts := user.Postlist(options)

	if posts == nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Unable to fetch post list for the specified user",
			Message:      "user.Postlist error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	out, err := SelectFields(posts, c.Request())
	if err == nil {
		return c.JSON(http.StatusOK, &Response{
			Data:         out,
			HumanMessage: "Correctly fetched post list for the specified user",
			Message:      "user.Postlist ok",
			Status:       http.StatusOK,
			Success:      true,
		})
	}
	return c.JSON(http.StatusBadRequest, &Response{
		HumanMessage: "Error selecting required fields",
		Message:      err.Error(),
		Status:       http.StatusBadRequest,
		Success:      false,
	})

}

/*
UserInfo handles the request and returns all the basic information for the
specified user
*/
func UserInfo(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	returnStruct := map[string]interface{}{
		"name":     user.Name,
		"contacts": user.ContactInfo(),
		"language": user.Language(),
		"email":    user.Email,
		"gender":   user.Gender,
		"personal": user.PersonalInfo(),
	}

	return c.JSON(http.StatusOK, &Response{
		HumanMessage: "Correctly retrieved user information",
		Data:         returnStruct,
		Message:      "User.Info ok",
		Status:       http.StatusOK,
		Success:      true,
	})

}

/*
UserFriends handles the request and returns the friend's of the specified user
*/
func UserFriends(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	users := user.Friends()
	usersStruct := map[string]interface{}{}

	for _, u := range *users {
		usersStruct[u.Username] = map[string]interface{}{
			"name":    u.Name,
			"surname": u.Surname,
			"from":    u.RegistrationTime,
		}
	}

	// Ops. No friends found
	if len(usersStruct) == 0 {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Unable to retrieve friends for the specified user",
			Message:      "User.Friends empty friends data",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	return c.JSON(http.StatusBadRequest, &Response{
		HumanMessage: "Correctly retrieved friends",
		Data:         usersStruct,
		Message:      "User.Friends ok",
		Status:       http.StatusOK,
		Success:      true,
	})

}
