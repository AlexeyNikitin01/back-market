// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OrdersProduct is an object representing the database table.
type OrdersProduct struct {
	OrderID           int64 `boil:"order_id" json:"orderID" toml:"orderID" yaml:"orderID"`
	ProductID         int64 `boil:"product_id" json:"productID" toml:"productID" yaml:"productID"`
	Discount          int   `boil:"discount" json:"discount" toml:"discount" yaml:"discount"`
	QuantityProduct   int64 `boil:"quantity_product" json:"quantityProduct" toml:"quantityProduct" yaml:"quantityProduct"`
	TotalProductPrice int64 `boil:"total_product_price" json:"totalProductPrice" toml:"totalProductPrice" yaml:"totalProductPrice"`

	R *ordersProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ordersProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrdersProductColumns = struct {
	OrderID           string
	ProductID         string
	Discount          string
	QuantityProduct   string
	TotalProductPrice string
}{
	OrderID:           "order_id",
	ProductID:         "product_id",
	Discount:          "discount",
	QuantityProduct:   "quantity_product",
	TotalProductPrice: "total_product_price",
}

var OrdersProductTableColumns = struct {
	OrderID           string
	ProductID         string
	Discount          string
	QuantityProduct   string
	TotalProductPrice string
}{
	OrderID:           "orders_products.order_id",
	ProductID:         "orders_products.product_id",
	Discount:          "orders_products.discount",
	QuantityProduct:   "orders_products.quantity_product",
	TotalProductPrice: "orders_products.total_product_price",
}

// Generated where

var OrdersProductWhere = struct {
	OrderID           whereHelperint64
	ProductID         whereHelperint64
	Discount          whereHelperint
	QuantityProduct   whereHelperint64
	TotalProductPrice whereHelperint64
}{
	OrderID:           whereHelperint64{field: "\"orders_products\".\"order_id\""},
	ProductID:         whereHelperint64{field: "\"orders_products\".\"product_id\""},
	Discount:          whereHelperint{field: "\"orders_products\".\"discount\""},
	QuantityProduct:   whereHelperint64{field: "\"orders_products\".\"quantity_product\""},
	TotalProductPrice: whereHelperint64{field: "\"orders_products\".\"total_product_price\""},
}

// OrdersProductRels is where relationship names are stored.
var OrdersProductRels = struct {
	Order   string
	Product string
}{
	Order:   "Order",
	Product: "Product",
}

// ordersProductR is where relationships are stored.
type ordersProductR struct {
	Order   *Order   `boil:"Order" json:"Order" toml:"Order" yaml:"Order"`
	Product *Product `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
}

// NewStruct creates a new relationship struct
func (*ordersProductR) NewStruct() *ordersProductR {
	return &ordersProductR{}
}

func (r *ordersProductR) GetOrder() *Order {
	if r == nil {
		return nil
	}
	return r.Order
}

func (r *ordersProductR) GetProduct() *Product {
	if r == nil {
		return nil
	}
	return r.Product
}

// ordersProductL is where Load methods for each relationship are stored.
type ordersProductL struct{}

var (
	ordersProductAllColumns            = []string{"order_id", "product_id", "discount", "quantity_product", "total_product_price"}
	ordersProductColumnsWithoutDefault = []string{"order_id", "product_id", "discount", "quantity_product", "total_product_price"}
	ordersProductColumnsWithDefault    = []string{}
	ordersProductPrimaryKeyColumns     = []string{"order_id", "product_id"}
	ordersProductGeneratedColumns      = []string{}
)

type (
	// OrdersProductSlice is an alias for a slice of pointers to OrdersProduct.
	// This should almost always be used instead of []OrdersProduct.
	OrdersProductSlice []*OrdersProduct
	// OrdersProductHook is the signature for custom OrdersProduct hook methods
	OrdersProductHook func(context.Context, boil.ContextExecutor, *OrdersProduct) error

	ordersProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ordersProductType                 = reflect.TypeOf(&OrdersProduct{})
	ordersProductMapping              = queries.MakeStructMapping(ordersProductType)
	ordersProductPrimaryKeyMapping, _ = queries.BindMapping(ordersProductType, ordersProductMapping, ordersProductPrimaryKeyColumns)
	ordersProductInsertCacheMut       sync.RWMutex
	ordersProductInsertCache          = make(map[string]insertCache)
	ordersProductUpdateCacheMut       sync.RWMutex
	ordersProductUpdateCache          = make(map[string]updateCache)
	ordersProductUpsertCacheMut       sync.RWMutex
	ordersProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ordersProductAfterSelectHooks []OrdersProductHook

var ordersProductBeforeInsertHooks []OrdersProductHook
var ordersProductAfterInsertHooks []OrdersProductHook

var ordersProductBeforeUpdateHooks []OrdersProductHook
var ordersProductAfterUpdateHooks []OrdersProductHook

var ordersProductBeforeDeleteHooks []OrdersProductHook
var ordersProductAfterDeleteHooks []OrdersProductHook

var ordersProductBeforeUpsertHooks []OrdersProductHook
var ordersProductAfterUpsertHooks []OrdersProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrdersProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrdersProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrdersProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrdersProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrdersProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrdersProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrdersProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrdersProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrdersProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ordersProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrdersProductHook registers your hook function for all future operations.
func AddOrdersProductHook(hookPoint boil.HookPoint, ordersProductHook OrdersProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ordersProductAfterSelectHooks = append(ordersProductAfterSelectHooks, ordersProductHook)
	case boil.BeforeInsertHook:
		ordersProductBeforeInsertHooks = append(ordersProductBeforeInsertHooks, ordersProductHook)
	case boil.AfterInsertHook:
		ordersProductAfterInsertHooks = append(ordersProductAfterInsertHooks, ordersProductHook)
	case boil.BeforeUpdateHook:
		ordersProductBeforeUpdateHooks = append(ordersProductBeforeUpdateHooks, ordersProductHook)
	case boil.AfterUpdateHook:
		ordersProductAfterUpdateHooks = append(ordersProductAfterUpdateHooks, ordersProductHook)
	case boil.BeforeDeleteHook:
		ordersProductBeforeDeleteHooks = append(ordersProductBeforeDeleteHooks, ordersProductHook)
	case boil.AfterDeleteHook:
		ordersProductAfterDeleteHooks = append(ordersProductAfterDeleteHooks, ordersProductHook)
	case boil.BeforeUpsertHook:
		ordersProductBeforeUpsertHooks = append(ordersProductBeforeUpsertHooks, ordersProductHook)
	case boil.AfterUpsertHook:
		ordersProductAfterUpsertHooks = append(ordersProductAfterUpsertHooks, ordersProductHook)
	}
}

// One returns a single ordersProduct record from the query.
func (q ordersProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrdersProduct, error) {
	o := &OrdersProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for orders_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrdersProduct records from the query.
func (q ordersProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrdersProductSlice, error) {
	var o []*OrdersProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to OrdersProduct slice")
	}

	if len(ordersProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrdersProduct records in the query.
func (q ordersProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count orders_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ordersProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if orders_products exists")
	}

	return count > 0, nil
}

// Order pointed to by the foreign key.
func (o *OrdersProduct) Order(mods ...qm.QueryMod) orderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	return Orders(queryMods...)
}

// Product pointed to by the foreign key.
func (o *OrdersProduct) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ordersProductL) LoadOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrdersProduct interface{}, mods queries.Applicator) error {
	var slice []*OrdersProduct
	var object *OrdersProduct

	if singular {
		var ok bool
		object, ok = maybeOrdersProduct.(*OrdersProduct)
		if !ok {
			object = new(OrdersProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrdersProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrdersProduct))
			}
		}
	} else {
		s, ok := maybeOrdersProduct.(*[]*OrdersProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrdersProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrdersProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ordersProductR{}
		}
		args = append(args, object.OrderID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ordersProductR{}
			}

			for _, a := range args {
				if a == obj.OrderID {
					continue Outer
				}
			}

			args = append(args, obj.OrderID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orders`),
		qm.WhereIn(`orders.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Order")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Order")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Order = foreign
		if foreign.R == nil {
			foreign.R = &orderR{}
		}
		foreign.R.OrdersProducts = append(foreign.R.OrdersProducts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrderID == foreign.ID {
				local.R.Order = foreign
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.OrdersProducts = append(foreign.R.OrdersProducts, local)
				break
			}
		}
	}

	return nil
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ordersProductL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrdersProduct interface{}, mods queries.Applicator) error {
	var slice []*OrdersProduct
	var object *OrdersProduct

	if singular {
		var ok bool
		object, ok = maybeOrdersProduct.(*OrdersProduct)
		if !ok {
			object = new(OrdersProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrdersProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrdersProduct))
			}
		}
	} else {
		s, ok := maybeOrdersProduct.(*[]*OrdersProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrdersProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrdersProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ordersProductR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ordersProductR{}
			}

			for _, a := range args {
				if a == obj.ProductID {
					continue Outer
				}
			}

			args = append(args, obj.ProductID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`products`),
		qm.WhereIn(`products.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for products")
	}

	if len(productAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Product = foreign
		if foreign.R == nil {
			foreign.R = &productR{}
		}
		foreign.R.OrdersProducts = append(foreign.R.OrdersProducts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.OrdersProducts = append(foreign.R.OrdersProducts, local)
				break
			}
		}
	}

	return nil
}

// SetOrder of the ordersProduct to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.OrdersProducts.
func (o *OrdersProduct) SetOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Order) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"orders_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
		strmangle.WhereClause("\"", "\"", 2, ordersProductPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrderID, o.ProductID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.ID
	if o.R == nil {
		o.R = &ordersProductR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &orderR{
			OrdersProducts: OrdersProductSlice{o},
		}
	} else {
		related.R.OrdersProducts = append(related.R.OrdersProducts, o)
	}

	return nil
}

// SetProduct of the ordersProduct to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.OrdersProducts.
func (o *OrdersProduct) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"orders_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
		strmangle.WhereClause("\"", "\"", 2, ordersProductPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrderID, o.ProductID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProductID = related.ID
	if o.R == nil {
		o.R = &ordersProductR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			OrdersProducts: OrdersProductSlice{o},
		}
	} else {
		related.R.OrdersProducts = append(related.R.OrdersProducts, o)
	}

	return nil
}

// OrdersProducts retrieves all the records using an executor.
func OrdersProducts(mods ...qm.QueryMod) ordersProductQuery {
	mods = append(mods, qm.From("\"orders_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"orders_products\".*"})
	}

	return ordersProductQuery{q}
}

// FindOrdersProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrdersProduct(ctx context.Context, exec boil.ContextExecutor, orderID int64, productID int64, selectCols ...string) (*OrdersProduct, error) {
	ordersProductObj := &OrdersProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"orders_products\" where \"order_id\"=$1 AND \"product_id\"=$2", sel,
	)

	q := queries.Raw(query, orderID, productID)

	err := q.Bind(ctx, exec, ordersProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from orders_products")
	}

	if err = ordersProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ordersProductObj, err
	}

	return ordersProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrdersProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no orders_products provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ordersProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ordersProductInsertCacheMut.RLock()
	cache, cached := ordersProductInsertCache[key]
	ordersProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ordersProductAllColumns,
			ordersProductColumnsWithDefault,
			ordersProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ordersProductType, ordersProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ordersProductType, ordersProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"orders_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"orders_products\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into orders_products")
	}

	if !cached {
		ordersProductInsertCacheMut.Lock()
		ordersProductInsertCache[key] = cache
		ordersProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrdersProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrdersProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ordersProductUpdateCacheMut.RLock()
	cache, cached := ordersProductUpdateCache[key]
	ordersProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ordersProductAllColumns,
			ordersProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update orders_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"orders_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ordersProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ordersProductType, ordersProductMapping, append(wl, ordersProductPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update orders_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for orders_products")
	}

	if !cached {
		ordersProductUpdateCacheMut.Lock()
		ordersProductUpdateCache[key] = cache
		ordersProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ordersProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for orders_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for orders_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrdersProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ordersProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"orders_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ordersProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in ordersProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all ordersProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrdersProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no orders_products provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ordersProductColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ordersProductUpsertCacheMut.RLock()
	cache, cached := ordersProductUpsertCache[key]
	ordersProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			ordersProductAllColumns,
			ordersProductColumnsWithDefault,
			ordersProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			ordersProductAllColumns,
			ordersProductPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert orders_products, could not build update column list")
		}

		ret := strmangle.SetComplement(ordersProductAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(ordersProductPrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert orders_products, could not build conflict column list")
			}

			conflict = make([]string, len(ordersProductPrimaryKeyColumns))
			copy(conflict, ordersProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"orders_products\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(ordersProductType, ordersProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ordersProductType, ordersProductMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert orders_products")
	}

	if !cached {
		ordersProductUpsertCacheMut.Lock()
		ordersProductUpsertCache[key] = cache
		ordersProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrdersProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrdersProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no OrdersProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ordersProductPrimaryKeyMapping)
	sql := "DELETE FROM \"orders_products\" WHERE \"order_id\"=$1 AND \"product_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from orders_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for orders_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ordersProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no ordersProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from orders_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for orders_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrdersProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ordersProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ordersProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"orders_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ordersProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from ordersProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for orders_products")
	}

	if len(ordersProductAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrdersProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrdersProduct(ctx, exec, o.OrderID, o.ProductID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrdersProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrdersProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ordersProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"orders_products\".* FROM \"orders_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ordersProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in OrdersProductSlice")
	}

	*o = slice

	return nil
}

// OrdersProductExists checks if the OrdersProduct row exists.
func OrdersProductExists(ctx context.Context, exec boil.ContextExecutor, orderID int64, productID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"orders_products\" where \"order_id\"=$1 AND \"product_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, orderID, productID)
	}
	row := exec.QueryRowContext(ctx, sql, orderID, productID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if orders_products exists")
	}

	return exists, nil
}

// Exists checks if the OrdersProduct row exists.
func (o *OrdersProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrdersProductExists(ctx, exec, o.OrderID, o.ProductID)
}
