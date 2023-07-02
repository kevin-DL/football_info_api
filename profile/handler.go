package profile

import (
	"encoding/json"
	"fmt"
	"football_api/ent"
	"football_api/ent/profile"
	"football_api/helpers"
	"github.com/google/uuid"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"log"
	"net/http"
)

type Controller struct {
	Client *ent.Client
}

func (receiver Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request CreateProfileRequest
	err := decoder.Decode(&request)

	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	// retrieve the session object as shown below
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	userID := sessionContainer.GetUserID()
	uid := uuid.MustParse(userID)
	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("failed to get uid: %s", err))
		return
	}

	u, err := receiver.Client.Profile.
		Create().
		SetDisplayName(request.Name).
		SetUserID(uid).
		Save(r.Context())
	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("failed creating profile: %s", err))
		return
	}
	log.Println("profile was created: ", u)
	//pro
	//return u, nil
	helpers.RespondWithJson(w, 201, u)
}

func (receiver Controller) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request CreateProfileRequest
	err := decoder.Decode(&request)

	if err != nil {
		helpers.RespondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	// retrieve the session object as shown below
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	userID := sessionContainer.GetUserID()
	uid := uuid.MustParse(userID)
	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("failed to get uid: %s", err))
		return
	}

	u, err := receiver.
		Client.
		Profile.
		Update().
		Where(profile.UserIDEQ(uid)).
		SetDisplayName(request.Name).
		Save(r.Context())

	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("failed creating profile: %s", err))
		return
	}
	log.Println("profile was updated: ", u)
	//pro
	//return u, nil
	helpers.RespondWithJson(w, 200, u)
}

func (receiver Controller) CurrentProfile(w http.ResponseWriter, r *http.Request) {
	// retrieve the session object as shown below
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	userID := sessionContainer.GetUserID()
	uid := uuid.MustParse(userID)

	u, err := receiver.
		Client.
		Profile.
		Query().
		Where(profile.UserIDEQ(uid)).
		First(r.Context())

	if err != nil {
		helpers.RespondWithError(w, 500, fmt.Sprintf("failed retrieving profile: %s", err))
		return
	}

	helpers.RespondWithJson(w, 200, u)
}
