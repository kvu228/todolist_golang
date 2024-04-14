package command

import (
	"context"
	"to_do_list/module/users/usecase"
)

type ChangeAvatarUseCase interface {
	ChangeAvatar(ctx context.Context, dto usecase.SetSingleImageDTO) error
}

type changeAvatarUseCase struct {
	userRepository  UserRepository
	imageRepository ImageRepository
}

func NewChangeAvatarUseCase(userRepository UserRepository, imageRepository ImageRepository) ChangeAvatarUseCase {
	return &changeAvatarUseCase{
		userRepository:  userRepository,
		imageRepository: imageRepository}
}

func (c *changeAvatarUseCase) ChangeAvatar(ctx context.Context, dto usecase.SetSingleImageDTO) error {
	userEntity, err := c.userRepository.FindById(ctx, dto.Requester.Id())
	if err != nil {
		return err
	}

	image, err := c.imageRepository.Find(ctx, dto.ImageId)
	if err != nil {
		return err
	}

	if err := userEntity.ChangeAvatar(image.FileName); err != nil {
		return err
	}

	if err := c.userRepository.Update(ctx, userEntity); err != nil {
		return err
	}

	//go func() {
	//	defer common.Recover()
	//
	//	ps := ctx.Value(common.CtxWithPubSub).(pubsub.PubSub)
	//	if err := ps.Publish(ctx, common.ChannelUserChangedAvater, pubsub.NewMessage(
	//		map[string]interface{}{
	//			"user_id": dto.Requester.UserId().String(),
	//			"img_id":  dto.ImageId.String(),
	//		})); err != nil {
	//		log.Println(err)
	//	}
	//
	//	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	//	return uc.imageRepo.SetImageStatusActivated(ctx, dto.ImageId)
	//	//}, asyncjob.WithName("SetImageStatus"))
	//	//
	//	//asyncjob.NewGroup(true, job).Run(ctx)
	//
	//}()

	return nil
}
