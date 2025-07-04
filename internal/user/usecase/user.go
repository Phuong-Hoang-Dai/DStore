package usecase

import "github.com/Phuong-Hoang-Dai/DStore/internal/user"

func CreateUser(uD UserCreateDTO, repos UserRepos) (int, error) {
	var u user.User
	MapUserCreateDTOtoUser(uD, &u)
	u.Role = "user"

	if id, err := repos.CreateUser(&u); err != nil {
		return 0, nil
	} else {
		return id, err
	}
}

func GetUserById(id int, repos UserRepos) (u UserResponeDTO, err error) {
	var userDAO user.User

	if err := repos.GetUserById(id, &userDAO); err != nil {
		return u, err
	} else {
		MapUsertoUserResponeDTo(userDAO, &u)
		return u, nil
	}
}

func UpdateUser(uD UserUpdateDTO, repos UserRepos) error {
	var u user.User
	MapUserUpdateDTOtoUser(uD, &u)

	if err := repos.UpdateUser(u); err != nil {
		return nil
	} else {
		return err
	}
}

func DeleteUser(id int, repos UserRepos) error {
	if err := repos.DeleteUser(id); err != nil {
		return err
	} else {
		return nil
	}
}

func GetUsers(p *user.Paging, repos UserRepos) ([]UserResponeDTO, error) {
	p.Process()
	var userDAO []user.User
	if err := repos.GetUsers(*p, &userDAO); err != nil {
		return nil, err
	}

	data := make([]UserResponeDTO, len(userDAO), cap(userDAO))
	for i := range userDAO {
		MapUsertoUserResponeDTo(userDAO[i], &data[i])
	}

	return data, nil
}

func Login(data UserDTO, repos UserRepos) (string, error) {
	userDao := user.User{}
	if err := repos.GetUserByName(data.Name, &userDao); err != nil {
		return "", err
	}

	if err := userDao.VerifyPassword([]byte(data.Password)); err != nil {
		return "", user.ErrUserNameOrPasswordIncorrect
	}

	if token, err := GenerateJwt(userDao); err != nil {
		return "", err
	} else {
		return token, err
	}
}
