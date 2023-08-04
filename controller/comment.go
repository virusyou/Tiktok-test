package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *fiber.Ctx) error {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			return c.Status(http.StatusOK).JSON(CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})

		}
		return c.Status(http.StatusOK).JSON(Response{StatusCode: 0})
	} else {
		return c.Status(http.StatusOK).JSON(Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})

}

// CommentAction no practical effect, just check if token is valid
//func CommentAction(c *gin.Context) {
//	token := c.Query("token")
//	actionType := c.Query("action_type")
//
//	if user, exist := usersLoginInfo[token]; exist {
//		if actionType == "1" {
//			text := c.Query("comment_text")
//			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
//				Comment: Comment{
//					Id:         1,
//					User:       user,
//					Content:    text,
//					CreateDate: "05-01",
//				}})
//			return
//		}
//		c.JSON(http.StatusOK, Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//// CommentList all videos have same demo comment list
//func CommentList(c *gin.Context) {
//	c.JSON(http.StatusOK, CommentListResponse{
//		Response:    Response{StatusCode: 0},
//		CommentList: DemoComments,
//	})
//}
