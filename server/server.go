package server

import (
	"context"
	"fmt"
	"log"

	"agate-project/db"
	"agate-project/handlers"
	"agate-project/repositories"
	"agate-project/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Server struct {
	DB                      *sqlx.DB
	Router                  *gin.Engine
	ClientRepo              repositories.ClientRepository
	ClientService           services.ClientService
	ClientHandlers          handlers.ClientHandlers
	StaffRepo               repositories.StaffRepository
	StaffService            services.StaffService
	StaffHandlers           handlers.StaffHandlers
	StaffGradeRepo          repositories.StaffGradeRepository
	StaffGradeService       services.StaffGradeService
	StaffGradeHandlers      handlers.StaffGradeHandlers
	CampaignRepo            repositories.CampaignRepository
	CampaignService         services.CampaignService
	CampaignHandlers        handlers.CampaignHandlers
	CampaignManagerRepo     repositories.CampaignManagerRepository
	CampaignManagerService  services.CampaignManagerService
	CampaignManagerHandlers handlers.CampaignManagerHandlers
	AdvertRepo              repositories.AdvertRepository
	AdvertService           services.AdvertService
	AdvertHandlers          handlers.AdvertHandlers
}

func NewServer(ctx context.Context) (*Server, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf(".env file could not be loaded: %v", err)
	}

	if err := db.OpenDatabase(); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	sqlxDB := sqlx.NewDb(db.DB, "postgres")

	clientRepo := repositories.NewClientRepository(ctx, sqlxDB)
	clientService := services.NewClientService(clientRepo)
	clientHandlers := handlers.NewClientHandlers(ctx, clientService)

	staffRepo := repositories.NewStaffRepository(ctx, sqlxDB)
	staffService := services.NewStaffService(staffRepo)
	staffHandlers := handlers.NewStaffHandlers(ctx, staffService)

	staffGradeRepo := repositories.NewStaffGradeRepository(ctx, sqlxDB)
	staffGradeService := services.NewStaffGradeService(staffGradeRepo)
	staffGradeHandlers := handlers.NewStaffGradeHandlers(ctx, staffGradeService)

	campaignRepo := repositories.NewCampaignRepository(ctx, sqlxDB)
	campaignService := services.NewCampaignService(campaignRepo)
	campaignHandlers := handlers.NewCampaignHandlers(campaignService)

	campaignManagerRepo := repositories.NewCampaignManagerRepository(ctx, sqlxDB)
	campaignManagerService := services.NewCampaignManagerService(campaignManagerRepo)
	campaignManagerHandlers := handlers.NewCampaignManagerHandlers(ctx, campaignManagerService)

	advertRepo := repositories.NewAdvertRepository(ctx, sqlxDB)
	advertService := services.NewAdvertService(advertRepo)
	advertHandlers := handlers.NewAdvertHandlers(ctx, advertService)

	router := gin.Default()

	router.GET("/clients", clientHandlers.GetClients)
	router.GET("/clients/:id", clientHandlers.GetClientByID)
	router.POST("/clients", clientHandlers.CreateClient)
	router.DELETE("/clients/:id", clientHandlers.RemoveClient)
	router.PUT("/clients/:id", clientHandlers.UpdateClient)

	router.GET("/staff", staffHandlers.GetStaff)
	router.GET("/staff/:id", staffHandlers.GetStaffByID)
	router.POST("/staff", staffHandlers.CreateStaff)
	router.DELETE("/staff/:id", staffHandlers.RemoveStaff)
	router.PUT("/staff/:id", staffHandlers.UpdateStaff)

	router.GET("/grades", staffGradeHandlers.GetAllGrades)
	router.POST("/grades", staffGradeHandlers.CreateGrade)
	router.DELETE("/grades/:id", staffGradeHandlers.RemoveGrade)
	router.PUT("/grades/:id", staffGradeHandlers.UpdateGrade)

	router.GET("/campaigns", campaignHandlers.GetAllCampaigns)
	router.POST("/campaigns", campaignHandlers.CreateCampaign)
	router.GET("/campaigns/:id", campaignHandlers.GetCampaignByID)
	router.PUT("/campaigns/:id", campaignHandlers.UpdateCampaign)
	router.DELETE("/campaigns/:id", campaignHandlers.RemoveCampaign)
	//router.GET("/campaigns/:id/budget", campaignHandlers.CheckBudget)
	router.PUT("/campaigns/:id/manager/:managerID", campaignHandlers.AssignManager)

	router.GET("/campaigns/client/:clientID", campaignHandlers.GetCampaignsByClientID)

	router.GET("/campaign-manager", campaignManagerHandlers.GetAllManagers)
	router.POST("/campaign-manager", campaignManagerHandlers.CreateManager)
	router.DELETE("/campaign-manager/:id", campaignManagerHandlers.DeleteManager)
	//router.PUT("/campaign-managers/:id", campaignManagerHandlers.UpdateManager)
	//router.POST("/campaign-managers/assign", campaignManagerHandlers.AssignStaffToCampaign)

	router.GET("/adverts", advertHandlers.GetAllAdverts)
	router.GET("/adverts/:id", advertHandlers.GetAdvertByID)
	router.POST("/adverts", advertHandlers.CreateAdvert)
	router.DELETE("/adverts/:id", advertHandlers.RemoveAdvert)
	router.PUT("/adverts/:id", advertHandlers.UpdateAdvert)
	router.GET("/adverts/campaign/:campaignID", advertHandlers.GetAdvertsByCampaign)

	srv := &Server{
		DB:                      sqlxDB,
		Router:                  router,
		ClientRepo:              clientRepo,
		ClientService:           clientService,
		ClientHandlers:          clientHandlers,
		StaffRepo:               staffRepo,
		StaffService:            staffService,
		StaffHandlers:           staffHandlers,
		StaffGradeRepo:          staffGradeRepo,
		StaffGradeService:       staffGradeService,
		StaffGradeHandlers:      staffGradeHandlers,
		CampaignRepo:            campaignRepo,
		CampaignService:         campaignService,
		CampaignHandlers:        campaignHandlers,
		CampaignManagerRepo:     campaignManagerRepo,
		CampaignManagerService:  campaignManagerService,
		CampaignManagerHandlers: campaignManagerHandlers,
		AdvertRepo:              advertRepo,
		AdvertService:           advertService,
		AdvertHandlers:          advertHandlers,
	}

	return srv, nil
}

func (s *Server) Run(addr string) error {
	return s.Router.Run(addr)
}

func (s *Server) Close() {
	db.CloseDatabase()
}
