package authmiddleware

import (
	// "errors"
	"log"
	"net/http"

	ijwt "bail/usecases/core/i_jwt"

	"github.com/gin-gonic/gin"
)

// Constants for context keys used in Gin middleware.
const (
	// ContextUserClaims is the key used to store user claims in the Gin context.
	ContextUserClaims = "userClaims"
)

// Authorize returns a Gin middleware handler that performs authentication and
// optional Authorizeation based on the provided JWT service and admin status requirement.
//
// It extracts the JWT from the "accessToken" cookie, decodes it, and checks if the user
// has the required admin status. If the user is authenticated and meets the Authorizeation
// criteria, their claims are attached to the request context; otherwise, an appropriate
// HTTP status code is returned and the request is aborted.
func AuthorizeAdmin(jwtService ijwt.Service, hasToBeAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from the cookie.
		cookie := c.GetHeader("accessToken")
		if cookie == "" {
			log.Println("no token found")
			c.Status(http.StatusBadRequest)
			c.Abort()
			return
		}
		// if err != nil {
		// 	if errors.Is(err, http.ErrNoCookie) {
		// 		log.Println("No cookie found")
		// 		c.Status(http.StatusUnauthorized) // No cookie found.
		// 	} else {
		// 		log.Println("Internal server error")
		// 		c.Status(http.StatusInternalServerError) // Internal server error.
		// 	}
		// 	c.Abort()
		// 	return
		// }

		// Decode the token using the JWT service.
		claims, err := jwtService.Decode(cookie)
		if err != nil {
			log.Println("Invalid token", claims, err)
			c.Status(http.StatusUnauthorized) // Invalid token.
			c.Abort()
			return
		}

		// Check if the user meets the required admin status.
		isAdmin, _ := claims["is_admin"].(bool)

		// if !ok{
		// 	log.Println("Internal server error")
		// 	c.Status(http.StatusInternalServerError) // Internal server error.
		// 	c.Abort()
		// 	return
		// }
		if !isAdmin && hasToBeAdmin {
			log.Println("Forbidden", isAdmin, hasToBeAdmin)
			c.Status(http.StatusForbidden) // Forbidden if admin status does not match.
			c.Abort()
			return
		}

		// Attach user claims to the request context for further use.
		c.Set(ContextUserClaims, claims)
		c.Next()
	}
}




func AuthorizeHR(jwtService ijwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from the cookie.
		cookie := c.GetHeader("accessToken")
		if cookie == "" {
			log.Println("no token found")
			c.Status(http.StatusBadRequest)
			c.Abort()
			return
		}
		// if err != nil {
		// 	if errors.Is(err, http.ErrNoCookie) {
		// 		log.Println("No cookie found")
		// 		c.Status(http.StatusUnauthorized) // No cookie found.
		// 	} else {
		// 		log.Println("Internal server error")
		// 		c.Status(http.StatusInternalServerError) // Internal server error.
		// 	}
		// 	c.Abort()
		// 	return
		// }

		// Decode the token using the JWT service.
		claims, err := jwtService.Decode(cookie)
		if err != nil {
			log.Println("Invalid token", claims, err)
			c.Status(http.StatusUnauthorized) // Invalid token.
			c.Abort()
			return
		}

		// Check if the user meets the required admin status.
		isHR, _ := claims["role"].(string)

		// if !ok{
		// 	log.Println("Internal server error")
		// 	c.Status(http.StatusInternalServerError) // Internal server error.
		// 	c.Abort()
		// 	return
		// }
		if isHR != "HR" && isHR != "Admin" {
			log.Println("Forbidden", isHR)
			c.Status(http.StatusForbidden) // Forbidden if admin status does not match.
			c.Abort()
			return
		}

		// Attach user claims to the request context for further use.
		c.Set(ContextUserClaims, claims)
		c.Next()
	}
}


func AuthorizeManager(jwtService ijwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from the cookie.
		cookie := c.GetHeader("accessToken")
		if cookie == "" {
			log.Println("no token found")
			c.Status(http.StatusBadRequest)
			c.Abort()
			return
		}
		// if err != nil {
		// 	if errors.Is(err, http.ErrNoCookie) {
		// 		log.Println("No cookie found")
		// 		c.Status(http.StatusUnauthorized) // No cookie found.
		// 	} else {
		// 		log.Println("Internal server error")
		// 		c.Status(http.StatusInternalServerError) // Internal server error.
		// 	}
		// 	c.Abort()
		// 	return
		// }

		// Decode the token using the JWT service.
		claims, err := jwtService.Decode(cookie)
		if err != nil {
			log.Println("Invalid token", claims, err)
			c.Status(http.StatusUnauthorized) // Invalid token.
			c.Abort()
			return
		}

		// Check if the user meets the required admin status.
		isManager, _ := claims["role"].(string)

		// if !ok{
		// 	log.Println("Internal server error")
		// 	c.Status(http.StatusInternalServerError) // Internal server error.
		// 	c.Abort()
		// 	return
		// }
		if isManager != "manager" && isManager != "Admin" {
			log.Println("Forbidden", isManager)
			c.Status(http.StatusForbidden) // Forbidden if admin status does not match.
			c.Abort()
			return
		}

		// Attach user claims to the request context for further use.
		c.Set(ContextUserClaims, claims)
		c.Next()
	}
}


func Authorize(jwtService ijwt.Service, hasToBeAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from the cookie.
		cookie := c.GetHeader("accessToken")
		if cookie == "" {
			log.Println("no token found")
			c.Status(http.StatusBadRequest)
			c.Abort()
			return
		}
		// if err != nil {
		// 	if errors.Is(err, http.ErrNoCookie) {
		// 		log.Println("No cookie found")
		// 		c.Status(http.StatusUnauthorized) // No cookie found.
		// 	} else {
		// 		log.Println("Internal server error")
		// 		c.Status(http.StatusInternalServerError) // Internal server error.
		// 	}
		// 	c.Abort()
		// 	return
		// }

		// Decode the token using the JWT service.
		claims, err := jwtService.Decode(cookie)
		if err != nil {
			log.Println("Invalid token", claims, err)
			c.Status(http.StatusUnauthorized) // Invalid token.
			c.Abort()
			return
		}

		// Check if the user meets the required admin status.
		

		// if !ok{
		// 	log.Println("Internal server error")
		// 	c.Status(http.StatusInternalServerError) // Internal server error.
		// 	c.Abort()
		// 	return
		// }

		// Attach user claims to the request context for further use.
		c.Set(ContextUserClaims, claims)
		c.Next()
	}
}