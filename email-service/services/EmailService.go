package services

import (
	email_pb "WxApp.kindle/email-service/proto/email"
	"context"
	"WxApp.kindle/email-service/models"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/RichardKnop/uuid"

)

//邮件服务数据结构，实现EmailServiceServer
type EmailService struct {
	db *gorm.DB
}

func (s *EmailService)CreateEmail(context.Context, req *email_pb.CreateEmailRequest) (*email_pb.CreateEmailResponse, error){
	user_id = req.GetUserId()
	if user_id == '' {
		return nil,error.New("UserId Not Found.")
	}
	email = req.GetEmail()
	if emial == '' {
		return nil,error.New("Email Not Found.")
	}
	user_email = models.UserEmail{
		BaseModel:{
			ID:			uuid.New(),
			CreatedAt: 	time.Time(),
			UpdatedAt: 	time.Time()
		},
		UserId:user_id,
		Email:email,
	}
	if err := s.db.Create(user_email).Error; err != nil {
		return nil, err
	}
	rep := &email_pb.CreateEmailResponse{
		Code:	1,
		Msg:	"添加成功"
	}
	return rep,nil
}
