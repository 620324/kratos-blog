package service

import (
	"context"
	pb "kratos-blog/api/blog"
	"kratos-blog/app/blog/internal/biz"
)

type BlogService struct {
	pb.UnimplementedBlogServer
	uc *biz.BlogUseCase
}

func NewBlogService(uc *biz.BlogUseCase) *BlogService {
	return &BlogService{uc: uc}
}

func (s *BlogService) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return s.uc.AddBlog(ctx, req), nil
}
func (s *BlogService) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogReply, error) {
	return s.uc.UpdateBlog(ctx, req), nil
}
func (s *BlogService) UpdateIndividualFields(ctx context.Context, req *pb.UpdateIndividualFieldsRequest) (*pb.UpdateIndividualFieldsReply, error) {
	return s.uc.UpdateIndividualFields(ctx, req), nil
}
func (s *BlogService) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogReply, error) {
	return s.uc.DeleteBlog(ctx, req), nil
}
func (s *BlogService) GetBlogByTag(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogReply, error) {
	return s.uc.GetByTagName(ctx, req), nil
}
func (s *BlogService) ListBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListBlogReply, error) {
	return s.uc.ListBlog(ctx, req), nil
}
func (s *BlogService) GetBlogByID(ctx context.Context, req *pb.GetBlogIDRequest) (*pb.GetBlogIDReply, error) {
	return s.uc.QueryBlogByID(ctx, req), nil
}
func (s *BlogService) GetBlogByTitle(ctx context.Context, req *pb.GetBlogByTitleRequest) (*pb.GetBlogByTitleReply, error) {
	return s.uc.QueryBlogByTitle(ctx, req), nil
}
func (s *BlogService) UpdateOnly(ctx context.Context, req *pb.UpdateOnlyRequest) (*pb.UpdateOnlyReply, error) {
	return s.uc.UpdateOnly(ctx, req), nil
}
func (s *BlogService) CacheBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return s.uc.CacheBlog(ctx, req), nil
}
func (s *BlogService) GetCacheBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListCacheReply, error) {
	return s.uc.GetAllCacheBlog(ctx), nil
}
func (s *BlogService) DeleteCacheBlog(ctx context.Context, req *pb.DeleteCacheBlogRequest) (*pb.DeleteCacheBlogReply, error) {
	return s.uc.DeleteCacheBlog(ctx, req), nil
}
func (s *BlogService) AddSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return s.uc.SetSuggestBlog(ctx, req), nil
}
func (s *BlogService) GetAllSuggest(ctx context.Context, req *pb.SearchAllSuggest) (*pb.SearchAllReply, error) {
	return s.uc.GetSuggestBlog(ctx, req), nil
}
func (s *BlogService) DeleteSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return s.uc.DeleteSuggestBlog(ctx, req), nil
}
func (s *BlogService) UpdateBlogVisit(ctx context.Context, request *pb.UpdateBlogVisitRq) (*pb.UpdateBlogVisitRq, error) {
	return s.uc.UpdateBlogVisit(ctx, request), nil
}
func (s *BlogService) GetCommentPower(ctx context.Context, req *pb.GetCommentPowerRq) (*pb.GetCommentPowerReply, error) {
	return s.uc.GetCommentPower(ctx, req), nil
}
