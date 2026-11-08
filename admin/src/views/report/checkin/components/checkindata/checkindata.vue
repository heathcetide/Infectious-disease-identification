<template>
  <el-drawer
    :model-value="visible"
    direction="rtl"
    size="800px"
    :with-header="false"
    custom-class="custom-drawer"
    @close="handleCancel"
  >
    <div class="container">
      <div class="title">
        <h3>职工打卡信息详情</h3>
      </div>
      <el-form
        :model="form"
        label-width="100px"
        class="form-container"
        ref="form"
        :rules="rules"
        disabled
      >
        <el-alert
          v-if="hybridDisagreement"
          type="warning"
          show-icon
          :closable="false"
          class="hybrid-alert"
          title="AI 与专家权重结论不一致"
          :description="hybridMessage || '请疾控人员复核 AI 证据链与专家 Top 结果后再研判。'"
        />
        <div v-if="aiScorePreview.length" class="ai-score-box">
          <div class="title-container">
            <div class="blue-box"></div>
            <span class="title-text">AI 风险摘要</span>
          </div>
          <div v-for="(row, idx) in aiScorePreview" :key="idx" class="ai-score-row">
            <span class="ai-name">{{ idx + 1 }}. {{ row.diseaseTypeName }}</span>
            <span class="ai-score">{{ Math.round(row.score || 0) }}</span>
            <span v-if="row.reason" class="ai-reason">{{ row.reason }}</span>
          </div>
        </div>
        <div class="BaseInfo">
          <div class="title-container">
            <div class="blue-box"></div>
            <span class="title-text">个人基本信息</span>
          </div>
          <div class="BaseInfoDetail">
            <!------------------------------- 姓名 --------------------------------------->
            <el-form-item label="姓名" prop="name">
              <el-input
                v-model="form.name"
                style="width: 200px"
                placeholder="请输入姓名"
                @blur="$refs.form.validateField('name')"
                clearable
              ></el-input>
            </el-form-item>

            <!------------------------------- 性别 ----------------------------------------->
            <el-form-item label="性别" prop="gender">
              <el-input
                v-model="form.gender"
                style="width: 200px"
                placeholder=""
                clearable
              ></el-input>
              <!-- <el-radio-group v-model="form.gender">
                <el-radio value="男">男</el-radio>
                <el-radio value="女">女</el-radio>
              </el-radio-group> -->
            </el-form-item>

            <!-- 年龄 -->
            <el-form-item label="年龄" prop="age">
              <el-input
                v-model="form.age"
                style="width: 200px"
                placeholder="根据身份证号生成"
                clearable
              ></el-input>
            </el-form-item>

            <!-- 民族 -->
            <el-form-item label="民族" prop="ethnicity">
              <el-input
                v-model="form.ethnicity"
                style="width: 200px"
                placeholder=""
                clearable
              ></el-input>
              <!-- <el-radio-group v-model="form.ethnicity">
                <el-radio value="汉族">汉族</el-radio>
                <el-radio value="藏族">藏族</el-radio>
                <el-radio value="彝族">彝族</el-radio>
                <el-radio value="其他少数民族">其他少数民族</el-radio>
              </el-radio-group> -->
            </el-form-item>

            <!-- 部门/工种 -->
            <el-form-item label="部门/工种" prop="department">
              <el-input
                v-model="form.department"
                style="width: 200px"
                placeholder=""
                clearable
              ></el-input>
              <!-- <el-radio-group v-model="form.department">
                <el-radio value="安全部">安全部</el-radio>
                <el-radio value="财务部">财务部</el-radio>
                <el-radio value="测量队">测量队</el-radio>
                <el-radio value="工程技术部">工程技术部</el-radio>
                <el-radio value="合约部">合约部</el-radio>
                <el-radio value="试验室">试验室</el-radio>
                <el-radio value="物资设备部">物资设备部</el-radio>
                <el-radio value="项目管理层">项目管理层</el-radio>
                <el-radio value="征拆协调部">征拆协调部</el-radio>
                <el-radio value="综合管理部">综合管理部</el-radio>
              </el-radio-group> -->
            </el-form-item>

            <el-form-item label="打卡时间" prop="checkInDate">
              <el-input
                v-model="form.checkInDate"
                style="width: 200px"
                @blur="$refs.form.validateField('checkInDate')"
                disabled
                clearable
              ></el-input>
            </el-form-item>
          </div>
        </div>

        <!------------------------------------ 症状标签 ------------------------------->
        <div class="select flex gap-2 mb-4">
          <div class="title-container">
            <div class="blue-box"></div>
            <span class="title-text">点击标签查看症状</span>
          </div>
          <el-check-tag
            :checked="selectedTag === 'GeneralSymptoms'"
            type="success"
            @change="selectTag('GeneralSymptoms')"
          >
            全身症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'RespiratorySymptoms'"
            type="primary"
            @change="selectTag('RespiratorySymptoms')"
          >
            呼吸系统症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'DigestiveSymptoms'"
            type="primary"
            @change="selectTag('DigestiveSymptoms')"
          >
            消化系统症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'CirculatorySymptoms'"
            type="primary"
            @change="selectTag('CirculatorySymptoms')"
          >
            循环系统症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'NeurologicalSymptoms'"
            type="primary"
            @change="selectTag('NeurologicalSymptoms')"
          >
            神经系统症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'LocalSymptoms'"
            type="primary"
            @change="selectTag('LocalSymptoms')"
          >
            局部症状
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'OtherSymptoms'"
            type="primary"
            @change="selectTag('OtherSymptoms')"
          >
            其他
          </el-check-tag>

          <el-check-tag
            :checked="selectedTag === 'RiskFactorsAndExposure'"
            type="primary"
            @change="selectTag('RiskFactorsAndExposure')"
          >
            危险因素与暴露史
          </el-check-tag>
        </div>

        <!-- 根据选择的标签显示不同内容 -->
        <div>
          <div v-show="selectedTag === 'GeneralSymptoms'">
            <p>查看全身症状</p>
            <GeneralSymptoms
              ref="GeneralSymptoms"
              :data="generalSymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'RespiratorySymptoms'">
            <p>查看呼吸系统症状</p>
            <RespiratorySymptoms
              ref="RespiratorySymptoms"
              :data="respiratorySymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'DigestiveSymptoms'">
            <p>查看消化系统症状</p>
            <DigestiveSymptoms
              ref="DigestiveSymptoms"
              :data="digestiveSymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'CirculatorySymptoms'">
            <p>查看循环系统症状</p>
            <CirculatorySymptoms
              ref="CirculatorySymptoms"
              :data="circulatorySymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'NeurologicalSymptoms'">
            <p>查看神经系统症状</p>
            <NeurologicalSymptoms
              ref="NeurologicalSymptoms"
              :data="neurologicalSymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'LocalSymptoms'">
            <p>查看局部症状</p>
            <DiagnosisLocalSymptoms
              ref="DiagnosisLocalSymptoms"
              :data="localSymptomsData"
            />
          </div>

          <div v-show="selectedTag === 'OtherSymptoms'">
            <p>查看其他症状</p>
            <OtherSymptoms ref="OtherSymptoms" :data="otherSymptomsData" />
          </div>

          <div v-show="selectedTag === 'RiskFactorsAndExposure'">
            <p>查看危险因素与暴露史</p>
            <RiskFactorsAndExposure
              ref="RiskFactorsAndExposure"
              :data="riskFactorsData"
            />
          </div>
        </div>
      </el-form>
    </div>
  </el-drawer>
</template>

<script>
import { ElMessage } from "element-plus";
import Dateselection from "@/components/date_selection.vue";
import GeneralSymptoms from "./GeneralSymptoms.vue";
import RespiratorySymptoms from "./RespiratorySymptoms.vue";
import CirculatorySymptoms from "./CirculatorySymptoms.vue";
import NeurologicalSymptoms from "./NeurologicalSymptoms.vue";
import DiagnosisLocalSymptoms from "./DiagnosisLocalSymptoms.vue";
import OtherSymptoms from "./OtherSymptoms.vue";
import RiskFactorsAndExposure from "./RiskFactorsAndExposure.vue";
import DigestiveSymptoms from "./DigestiveSymptoms.vue";
import { selectCheckinById } from "@/api/report/checkin.js";
export default {
  components: {
    Dateselection,
    GeneralSymptoms,
    RespiratorySymptoms,
    CirculatorySymptoms,
    NeurologicalSymptoms,
    DiagnosisLocalSymptoms,
    OtherSymptoms,
    RiskFactorsAndExposure,
    DigestiveSymptoms,
  },
  data() {
    return {
      visible: false, // 控制弹窗显示
      selectedTag: null, // 当前选中的标签
      form: {},
      rules: {},

      generalSymptomsData: null, // 全身症状数据
      respiratorySymptomsData: null, // 呼吸系统症状数据
      digestiveSymptomsData: null, // 消化系统症状数据
      circulatorySymptomsData: null, // 循环系统症状数据
      neurologicalSymptomsData: null, // 神经系统症状数据
      localSymptomsData: null, // 局部症状数据
      otherSymptomsData: null, // 其他症状数据
      hybridDisagreement: false,
      hybridMessage: '',
      aiScorePreview: [],
      riskFactorsData: null, // 危险因素与暴露史数据

      dailyHealthId: null,
    };
  },

  methods: {
    toggleTag(field) {
      this.form[field] = !this.form[field];
    },
    selectTag(tag) {
      this.selectedTag = tag;
    },
    showDrawer(user) {
      this.form = { ...user };
      this.visible = true;
      this.dailyHealthId = user.statusId;
      this.handleQuery(this.dailyHealthId);
    },
    handleCancel() {
      this.visible = false;
      this.handleReset();

      // if (this.$refs.GeneralSymptoms) {
      //   this.$refs.GeneralSymptoms.handleReset();
      // }
      // if (this.$refs.RespiratorySymptoms) {
      //   this.$refs.RespiratorySymptoms.handleReset();
      // }
      // if (this.$refs.CirculatorySymptoms) {
      //   this.$refs.CirculatorySymptoms.handleReset();
      // }
      // if (this.$refs.NeurologicalSymptoms) {
      //   this.$refs.NeurologicalSymptoms.handleReset();
      // }
      // if (this.$refs.DiagnosisLocalSymptoms) {
      //   this.$refs.DiagnosisLocalSymptoms.handleReset();
      // }
      // if (this.$refs.OtherSymptoms) {
      //   this.$refs.OtherSymptoms.handleReset();
      // }
      // if (this.$refs.RiskFactorsAndExposure) {
      //   this.$refs.RiskFactorsAndExposure.handleReset();
      // }
    },

 async handleQuery(dailyHealthId) {
      try {
        const response = await selectCheckinById(dailyHealthId);

        if (response && response.data && response.data.code === 1) {
          const data = response.data.data;

          console.log("完整的后端响应：", data);

          // 分发数据到相应的属性（兼容旧症状结构 + 新 CheckInDetailVO）
          this.generalSymptomsData = data.generalsymptoms || {};
          this.respiratorySymptomsData = data.respiratorysymptoms || {};
          this.digestiveSymptomsData = data.digestivesymptoms || {};
          this.circulatorySymptomsData = data.circulatorysymptoms || {};
          this.neurologicalSymptomsData = data.neurologicalSymptoms || {};
          this.localSymptomsData = data.localSymptoms || {};
          this.otherSymptomsData = data.otherSymptoms || {};
          this.riskFactorsData = data.riskFactorsAndExposure || {};

          const hybrid = data.hybrid || {};
          const status = data.status || {};
          this.hybridDisagreement = !!(
            hybrid.disagreement || status.scoringDisagreement
          );
          this.hybridMessage = hybrid.message || '';
          const aiList = Array.isArray(data.aiDiseaseScoringList)
            ? [...data.aiDiseaseScoringList]
            : [];
          this.aiScorePreview = aiList
            .sort((a, b) => (b.score || 0) - (a.score || 0))
            .slice(0, 5);

          console.log("分发后的循环系统数据:", this.circulatorySymptomsData);
        } else {
          console.error("获取数据失败或无效的响应:", response);
        }
      } catch (error) {
        console.error("获取打卡信息出错:", error);
        ElMessage.error("获取打卡信息失败");
      }
    },

    handleReset() {
      this.form = this.getInitialForm();
      this.message = "";
      this.selectedTag = null;
    },
    getInitialForm() {
      return {};
    },
  },
};
</script>



<style scoped>
.custom-drawer {
  height: 100%;
}
.form-container {
  padding: 20px;
  margin-top: 70px;
}
.hybrid-alert {
  margin-bottom: 14px;
}
.ai-score-box {
  margin-bottom: 18px;
  padding: 10px 12px;
  background: #f8fafc;
  border-radius: 8px;
}
.ai-score-row {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 4px 12px;
  padding: 6px 0;
  border-bottom: 1px solid #e2e8f0;
  font-size: 13px;
}
.ai-score-row:last-child {
  border-bottom: none;
}
.ai-name {
  font-weight: 600;
  color: #0f172a;
}
.ai-score {
  color: #dc2626;
  font-weight: 700;
}
.ai-reason {
  grid-column: 1 / -1;
  color: #64748b;
  font-size: 12px;
  line-height: 1.4;
}
.title {
  position: fixed;
  top: 0;
  right: 0;
  width: 800px;
  background: #ffffff;
  padding-bottom: 30px;
  z-index: 100;
  border: 3px solid #fafafa;
}

h3 {
  margin-top: 20px;
  font-size: 16px;
  color: #333333;
  margin-left: 30px;
}
.custom-input {
  margin-bottom: 10px;
}
.el-dropdown-link:focus {
  outline: none;
}

.title-container {
  display: flex;
  margin-left: 0px;
  margin-bottom: 20px;
  margin-top: 20px;
}

.blue-box {
  width: 6px;
  height: 18px;
  background-color: #285ac8;
  margin-right: 10px;
}

.title-text {
  font-size: 12px;
  font-weight: bold;
  color: #4a4a4a;
}

.BaseInfo {
  margin-bottom: 50px;
}
.el-check-tag {
  margin: 10px;
  font-weight: normal;
}
.el-check-tag.is-checked {
  background-color: #333333 !important;
  color: #ffffff;
  font-weight: normal;
}
.select {
  margin-top: 10px;
  margin-bottom: 20px;
}
</style>
