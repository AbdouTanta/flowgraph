package auth

type AuthService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (s *AuthService) Login(loginPayload User) (User, error) {
	user, err := s.authRepository.Login(&loginPayload)
	if err != nil {
		return User{}, err
	}
	return *user, err
}
