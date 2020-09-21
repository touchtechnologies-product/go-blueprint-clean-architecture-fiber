package service

import (
	"context"

	"github.com/touchtechnologies-product/goerror"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
)

var (
	ErrStaffNotFound   = goerror.DefineNotFound("StaffNotFound", "staff not found")
	ErrUnableGetStaff  = goerror.DefineNotFound("UnableGetStaff", "unable to get staff")
	ErrUnableGetStaffs = goerror.DefineNotFound("UnableGetStaffs", "unable to get staffs by company")
	ErrUnableSaveStaff = goerror.DefineNotFound("UnableSaveStaff", "unable to save staff")
)

//go:generate mockery -name=Repository
type Repository interface {
	Get(ctx context.Context, staffId string) (*staff.Staff, goerror.Error)
	GetStaffsByCompany(ctx context.Context, companyId string, offset, limit int64) ([]*staff.Staff, goerror.Error)
	Save(ctx context.Context, staff *staff.Staff) goerror.Error
}
