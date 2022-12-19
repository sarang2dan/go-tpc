package tpch

import (
	"context"
	"fmt"
	"github.com/pingcap/go-tpc/pkg/util"
	log "github.com/sirupsen/logrus"
	"strings"
)

var allTables []string

func init() {
	allTables = []string{"lineitem", "partsupp", "supplier", "part", "orders", "customer", "region", "nation"}
}

func (w *Workloader) createTableDDL(ctx context.Context, query string, tableName string, action string) error {
	s := w.getState(ctx)
	fmt.Printf("%s %s\n", action, tableName)
	if _, err := s.Conn.ExecContext(ctx, query); err != nil {
		return err
	}
	if w.cfg.TiFlashReplica != 0 {
		fmt.Printf("creating tiflash replica for %s\n", tableName)
		replicaSQL := fmt.Sprintf("ALTER TABLE %s SET TIFLASH REPLICA %d", tableName, w.cfg.TiFlashReplica)
		if _, err := s.Conn.ExecContext(ctx, replicaSQL); err != nil {
			return err
		}
	}
	return nil
}

// createTables creates tables schema.
func (w *Workloader) createTables(ctx context.Context) error {
	ddls, err := util.GetDDLQueries(w.cfg.Driver.String(), "tpch")
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	query := strings.Join(ddls.Nation, "\n")
	if err := w.createTableDDL(ctx, query, "nation", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Region, "\n")
	if err := w.createTableDDL(ctx, query, "region", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Part, "\n")
	if err := w.createTableDDL(ctx, query, "part", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Supplier, "\n")
	if err := w.createTableDDL(ctx, query, "supplier", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.PartSupp, "\n")
	if err := w.createTableDDL(ctx, query, "partsupp", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Customer, "\n")
	if err := w.createTableDDL(ctx, query, "customer", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Orders, "\n")
	if err := w.createTableDDL(ctx, query, "orders", "creating"); err != nil {
		return err
	}

	query = strings.Join(ddls.Lineitem, "\n")
	if err := w.createTableDDL(ctx, query, "lineitem", "creating"); err != nil {
		return err
	}
	return nil
}

func (w *Workloader) dropTable(ctx context.Context) error {
	s := w.getState(ctx)

	for _, tbl := range allTables {
		fmt.Printf("DROP TABLE IF EXISTS %s\n", tbl)
		if _, err := s.Conn.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", tbl)); err != nil {
			return err
		}
	}
	return nil
}
