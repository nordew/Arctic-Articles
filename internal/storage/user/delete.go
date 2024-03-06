package user

import "context"

func (s *userStorage) Delete(ctx context.Context, userID string) error {
	const op = "userStorage.Delete"

	query := "DELETE FROM users WHERE id = $1"

	_, err := s.conn.Exec(ctx, query, userID)
	if err != nil {
		s.logger.Error("failed to delete user", err.Error(), op)
		return err
	}

	return nil
}
