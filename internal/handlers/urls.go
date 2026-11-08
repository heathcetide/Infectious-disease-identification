// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/llm"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/apidocs/humax"
	"github.com/heathcetide/infectious-disease/pkg/common/jwtutil"
	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppInfo holds application metadata.
type AppInfo struct {
	Name      string
	Version   string
	BuildTime string
	GitCommit string
	JWTAuth *jwtutil.Auth
}

// Handlers is the HTTP handler collection.
// Storage/cache/lock/retry are accessed via pkg/ package-level singletons.
type Handlers struct {
	cfg       *configs.Config
	info      AppInfo
	db        *gorm.DB
	startTime time.Time
	jwtAuth   *jwtutil.Auth
	svc       *services.Services
}

// New creates a new Handlers instance.
func New(cfg *configs.Config, info AppInfo, db *gorm.DB) *Handlers {
	chat := llm.New(llm.Config{
		Enabled:  cfg.LLM.Enabled,
		Protocol: cfg.LLM.Protocol,
		APIKey:   cfg.LLM.APIKey,
		BaseURL:  cfg.LLM.BaseURL,
		Model:    cfg.LLM.Model,
		Timeout:  cfg.LLM.Timeout,
	})
	return &Handlers{
		cfg:       cfg,
		info:      info,
		db:        db,
		startTime: time.Now(),
		jwtAuth:   info.JWTAuth,
		svc:       services.New(db, chat),
	}
}

// Register registers all routes via humax.Group (Gin + OpenAPI docs).
func (h *Handlers) Register(engine *gin.Engine, api huma.API) {
	r := humax.NewGroup(api, engine, "")
	h.registerRoutes(r)
}

// registerRoutes registers all route groups.
func (h *Handlers) registerRoutes(r *humax.Group) {
	h.registerSystemRoutes(r)
	h.registerAuthRoutes(r)
	h.registerUserRoutes(r)
	h.registerFileRoutes(r)
	h.registerEmployeeRoutes(r)
	h.registerCheckInRoutes(r)
	h.registerDiagnosisRoutes(r)
	h.registerLabRoutes(r)
	h.registerLearningRoutes(r)
	h.registerInfectFileRoutes(r)
	h.registerAdminInfectRoutes(r)
	h.registerAdminSatisfactionRoutes(r)
	h.registerAdminLegacyRoutes(r)
	h.registerAIRoutes(r)
	h.registerCDCAccuracyRoutes(r)
	h.registerCDCRoutes(r)
	h.registerMedicalRoutes(r)
}

// registerSystemRoutes registers public system endpoints (health, version).
func (h *Handlers) registerSystemRoutes(r *humax.Group) {
	r.GET("/health", h.Health)
	r.GET("/live", h.Liveness)
	r.GET("/ready", h.Readiness)
	r.GET("/api/v1/version", h.Version)
}

// registerAuthRoutes registers public auth endpoints (login, refresh).
func (h *Handlers) registerAuthRoutes(r *humax.Group) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}
	r.POST("/user/login", h.UserLogin)
}

// registerUserRoutes registers user registration (public) + user CRUD endpoints.
func (h *Handlers) registerUserRoutes(r *humax.Group) {
	// Public registration
	r.POST("/api/v1/register", h.RegisterUser)

	// User CRUD
	users := r.Group("/api/v1/users")
	{
		users.GET("", h.ListUsers)
		users.GET("/:id", h.GetUser)
		users.POST("", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.PUT("/:id/password", h.ChangePassword)
		users.DELETE("/:id", h.DeleteUser)
	}
}

// registerFileRoutes registers file upload/download/delete endpoints.
func (h *Handlers) registerFileRoutes(r *humax.Group) {
	files := r.Group("/api/v1/files")
	{
		files.POST("/upload", h.UploadFile)
		files.GET("/:key", h.DownloadFile)
		files.DELETE("/:key", h.DeleteFile)
	}
}

func (h *Handlers) registerEmployeeRoutes(r *humax.Group) {
	emp := r.Group("/api/v1/employee")
	{
		emp.GET("/profile", h.EmployeeGetProfile)
		emp.PUT("/profile", h.EmployeeUpdateProfile)
		emp.PUT("/password", h.EmployeeChangePassword)
		emp.POST("/feedback", h.EmployeeSubmitFeedback)
		emp.GET("/satisfaction/need-rating", h.EmployeeNeedRating)
		emp.POST("/satisfaction", h.EmployeeSubmitSatisfaction)
		emp.GET("/hospitals/nearby", h.EmployeeNearbyHospitals)
	}
}

func (h *Handlers) registerCheckInRoutes(r *humax.Group) {
	checkin := r.Group("/api/v1/checkin")
	{
		checkin.POST("", h.CheckInGeo)
		checkin.GET("/daily", h.CheckInDaily)
		checkin.GET("/month", h.CheckInMonth)
		checkin.GET("/reverse-geocode", h.CheckInReverseGeocode)
		checkin.POST("/symptoms", h.CheckInSymptoms)
		checkin.POST("/symptom-photo", h.CheckInUploadSymptomPhoto)
		checkin.POST("/scoring/confirm", h.CheckInConfirmScoring)
		checkin.POST("/environment", h.CheckInEnvironment)
		checkin.PUT("/environment", h.CheckInEnvironmentUpdate)
		checkin.GET("/environment", h.CheckInEnvironmentGet)
	}
}

func (h *Handlers) registerDiagnosisRoutes(r *humax.Group) {
	diag := r.Group("/api/v1/diagnosis")
	{
		diag.POST("/result", h.DiagnosisSubmitResult)
		diag.GET("/report", h.DiagnosisGetReport)
		diag.POST("/file", h.DiagnosisUploadFile)
	}
}

func (h *Handlers) registerLabRoutes(r *humax.Group) {
	lab := r.Group("/api/v1/lab")
	{
		lab.POST("/report", h.LabSubmitReport)
		lab.GET("/report", h.LabGetReport)
		lab.POST("/file", h.LabUploadFile)
	}
}

func (h *Handlers) registerLearningRoutes(r *humax.Group) {
	learning := r.Group("/api/v1/learning")
	{
		learning.POST("/materials/page", h.LearningMaterialsPage)
		learning.POST("/materials/:id/study", h.LearningMaterialStudy)
	}
}

func (h *Handlers) registerInfectFileRoutes(r *humax.Group) {
	inf := r.Group("/api/v1/infect-files")
	{
		inf.GET("/diagnosis/:id", h.InfectDiagnosisFile)
		inf.GET("/lab/:id", h.InfectLabFile)
		inf.GET("/learning/:filename", h.InfectLearningFile)
	}
}

func (h *Handlers) registerAdminInfectRoutes(r *humax.Group) {
	admin := r.Group("/api/v1/admin")
	{
		admin.GET("/hello", h.AdminHello)
		admin.GET("/dashboard", h.AdminDashboard)
		admin.POST("/users/page", h.AdminUsersPage)
		admin.GET("/users", h.ListUsers)
		admin.GET("/users/export", h.AdminExportUsers)
		admin.POST("/users/import", h.AdminImportUsers)
		admin.GET("/users/:id", h.AdminGetUser)
		admin.POST("/users", h.AdminCreateStaffUser)
		admin.PUT("/users/:id", h.UpdateUser)
		admin.PUT("/users/:id/active", h.AdminSetUserActive)
		admin.PUT("/users/:id/password/reset", h.AdminResetUserPassword)
		admin.DELETE("/users/:id", h.DeleteUser)
		admin.POST("/hospitals/page", h.AdminHospitalPage)
		admin.GET("/hospitals", h.AdminListHospitals)
		admin.GET("/hospitals/export", h.AdminExportHospitals)
		admin.POST("/hospitals/import", h.AdminImportHospitals)
		admin.GET("/hospitals/import/template", h.AdminHospitalsImportTemplate)
		admin.GET("/hospitals/:id", h.AdminGetHospital)
		admin.POST("/hospitals", h.AdminCreateHospital)
		admin.PUT("/hospitals/:id", h.AdminUpdateHospital)
		admin.DELETE("/hospitals/:id", h.AdminDeleteHospital)
		admin.POST("/diseases/page", h.AdminDiseasesPage)
		admin.GET("/diseases", h.AdminListDiseases)
		admin.GET("/diseases/export", h.AdminExportDiseases)
		admin.GET("/weights", h.AdminListWeights)
		admin.PUT("/weights/batch", h.AdminWeightsBatch)
		admin.GET("/weights/export", h.AdminExportWeights)
		admin.GET("/weights/export-template", h.AdminWeightsExportTemplate)
		admin.POST("/weights/import", h.AdminImportWeights)
		admin.GET("/feedback", h.AdminListFeedback)
		admin.GET("/feedback/export", h.AdminExportFeedback)
		admin.GET("/feedback/:id", h.AdminGetFeedback)
		admin.GET("/settings/ai-mode", h.AdminGetAIMode)
		admin.PUT("/settings/ai-mode", h.AdminSetAIMode)
		admin.POST("/weights/batch-with-diagnosis", h.AdminWeightsBatchWithDiagnosis)
		admin.PUT("/weights/batch-with-diagnosis", h.AdminWeightsBatchWithDiagnosis)

		admin.GET("/llm/info", h.AdminLLMInfo)
		admin.POST("/llm/chat", h.AdminLLMChat)

		admin.GET("/vision/samples", h.AdminVisionSamplesPage)
		admin.POST("/vision/samples", h.AdminVisionSampleCreate)
		admin.GET("/vision/samples/:id/file", h.AdminVisionSampleFile)
		admin.DELETE("/vision/samples/:id", h.AdminVisionSampleDelete)
		admin.GET("/vision/summary", h.AdminVisionDatasetSummary)
		admin.GET("/vision/export", h.AdminVisionDatasetExport)
		admin.POST("/vision/classify", h.AdminVisionClassifyTrial)
		admin.POST("/vision/promote", h.AdminVisionPromotePhoto)

		admin.GET("/eval/versions", h.AdminListScoringVersions)
		admin.POST("/eval/versions", h.AdminCreateScoringVersion)
		admin.GET("/eval/versions/:id", h.AdminGetScoringVersion)
		admin.GET("/eval/versions/:id/download", h.AdminDownloadScoringVersion)
		admin.POST("/eval/versions/:id/restore", h.AdminRestoreScoringVersion)
		admin.GET("/eval/runs", h.AdminListEvalRuns)
		admin.POST("/eval/runs", h.AdminRunEval)
		admin.GET("/eval/runs/:id/cases", h.AdminListEvalCases)

		admin.GET("/risk-alerts", h.AdminListRiskAlerts)
		admin.PUT("/risk-alerts/:id/ack", h.AdminAckRiskAlert)
	}
}

func (h *Handlers) registerAdminSatisfactionRoutes(r *humax.Group) {
	sat := r.Group("/api/v1/admin/satisfaction")
	{
		sat.POST("/surveys", h.AdminCreateSatisfactionSurvey)
		sat.PUT("/surveys/:id", h.AdminUpdateSatisfactionSurvey)
		sat.PUT("/surveys/:id/open", h.AdminOpenSatisfactionSurvey)
		sat.PUT("/surveys/:id/close", h.AdminCloseSatisfactionSurvey)
		sat.GET("/surveys/:id/stats", h.AdminSatisfactionStats)
		sat.GET("/surveys", h.AdminPageSatisfactionSurveys)
		sat.GET("/ratings", h.AdminPageSatisfactionRatings)
	}
}

func (h *Handlers) registerAdminLegacyRoutes(r *humax.Group) {
	legacy := r.Group("/admin/satisfaction")
	{
		legacy.POST("/addSatisfactionSurvey", h.AdminCreateSatisfactionSurveyLegacy)
		legacy.POST("/updateSatisfaction", h.AdminUpdateSatisfactionLegacy)
		legacy.GET("/closeSatisfaction/:surveyID", h.AdminCloseSatisfactionLegacy)
		legacy.GET("/openSatisfaction/:surveyID", h.AdminOpenSatisfactionLegacy)
		legacy.GET("/getNumberInfo/:surveyID", h.AdminSatisfactionStatsLegacy)
		legacy.GET("/pageSelectSatisfactionSurvey", h.AdminPageSatisfactionSurveysLegacy)
		legacy.GET("/pageSelectSatisfactionRating", h.AdminPageSatisfactionRatingsLegacy)
	}
}

func (h *Handlers) registerAIRoutes(r *humax.Group) {
	ai := r.Group("/api/v1/ai")
	{
		ai.GET("/auto-weight", h.AIGetAutoWeight)
		ai.PUT("/auto-weight", h.AISetAutoWeight)
		ai.PUT("/auto-weight/:enabled", h.AISetAutoWeight)
		ai.POST("/auto-weight/run", h.AIRunAutoWeight)
	}
	r.GET("/ai/getAutoAIWeightUpdater", h.AIGetAutoWeightLegacy)
	r.GET("/ai/updateAutoAIWeightUpdater/:autoAIWeightUpdater", h.AISetAutoWeightLegacy)
	r.GET("/ai/startAI", h.AIRunAutoWeightLegacy)
}

func (h *Handlers) registerCDCAccuracyRoutes(r *humax.Group) {
	acc := r.Group("/api/v1/cdc/accuracy")
	{
		acc.GET("/strike", h.CDCAccuracyStrike)
		acc.PUT("/predictor-range/:range", h.CDCSetPredictorRange)
		acc.GET("/predictor/:range", h.CDCGetPredictorRange)
	}
	r.GET("/cdcStaff/statisticsDiagnosticAccuracy/strike", h.CDCAccuracyStrikeLegacy)
	r.GET("/cdcStaff/statisticsDiagnosticAccuracy/JudgmentsPredictorNumber/:range", h.CDCSetPredictorRangeLegacy)
}

func (h *Handlers) registerCDCRoutes(r *humax.Group) {
	cdc := r.Group("/api/v1/cdc")
	{
		cdc.POST("/checkin/stats/page", h.CDCCheckInStatsPage)
		cdc.GET("/checkin/stats/export", h.CDCCheckInStatsExport)
		cdc.GET("/checkin/stats/:id/location", h.CDCCheckInStatsLocation)
		cdc.GET("/checkin/stats/:id", h.CDCCheckInStatsByID)
		cdc.GET("/dashboard", h.CDCDashboard)
		cdc.GET("/dashboard/statistics", h.CDCDashboardStatistics)
		cdc.GET("/dashboard/disease-today", h.CDCDashboardDiseaseToday)
		cdc.GET("/dashboard/user-stations", h.CDCDashboardUserStations)
		cdc.GET("/dashboard/checkin-search", h.CDCDashboardCheckinSearch)
		cdc.POST("/dashboard/checkin-daily-series", h.CDCDashboardCheckinDailySeries)
		cdc.GET("/dashboard/important-users", h.CDCDashboardImportantUsers)
		cdc.GET("/disease/stats", h.CDCDiseaseStats)
		cdc.GET("/disease/statistics", h.CDCDiseaseStatistics)
		cdc.GET("/work-environments", h.CDCWorkEnvironments)
		cdc.GET("/work-environments/export", h.CDCExportWorkEnvironments)
		cdc.GET("/work-environments/:id", h.CDCGetWorkEnvironment)
		cdc.GET("/materials", h.CDCMaterialsManager)
		cdc.POST("/materials", h.CDCCreateMaterial)
		cdc.GET("/materials/:id", h.CDCGetMaterial)
		cdc.PUT("/materials/:id", h.CDCUpdateMaterial)
		cdc.PUT("/materials/:id/status", h.CDCMaterialStatus)
		cdc.PUT("/materials/:id/sequence", h.CDCMaterialSequence)
		cdc.GET("/eval/latest", h.CDCLatestEvalCompare)
	}
}

func (h *Handlers) registerMedicalRoutes(r *humax.Group) {
	med := r.Group("/api/v1/medical")
	{
		med.POST("/diagnosis/page", h.MedicalDiagnosisManager)
		med.GET("/diagnosis/export", h.MedicalExportDiagnosis)
		med.GET("/diagnosis/:id", h.MedicalGetDiagnosis)
		med.POST("/diagnosis", h.MedicalSaveDiagnosis)
		med.PUT("/diagnosis/:id", h.MedicalUpdateDiagnosis)
		med.POST("/lab/page", h.MedicalLabManager)
		med.GET("/lab/export", h.MedicalExportLab)
		med.GET("/lab/:id", h.MedicalGetLab)
		med.POST("/lab", h.MedicalCreateLab)
		med.PUT("/lab/:id", h.MedicalUpdateLab)
		med.POST("/users/lookup", h.MedicalUserLookup)
		med.GET("/users/:id", h.MedicalGetUser)
		med.PUT("/password", h.MedicalChangePassword)
	}
}
