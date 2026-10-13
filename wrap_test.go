package modelg

import "testing"

func TestUnwrapFindsTheDatabase(t *testing.T) {
	pgx := NewPgxDB(nil)
	if found, ok := Unwrap[*PgxDB](NewErrorTranslatingDB(pgx, nil)); !ok || found != pgx {
		t.Errorf("Unwrap[*PgxDB] = %v, %v", found, ok)
	}
	sqlite := NewSqliteDB(nil)
	if found, ok := Unwrap[*SqliteDB](NewErrorTranslatingDB(sqlite, nil)); !ok || found != sqlite {
		t.Errorf("Unwrap[*SqliteDB] = %v, %v", found, ok)
	}
}
