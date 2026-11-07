<template>
  <el-drawer
    :model-value="visible"
    direction="rtl"
    size="700px"
    :with-header="false"
    custom-class="custom-drawer"
    @close="handleCancel"
  >
    <div class="container">
      <div class="title">
        <h3>职工打卡位置详情</h3>
      </div>
      <el-form
        :model="form"
        label-width="100px"
        class="form-container"
        ref="form"
        :rules="rules"
        disabled
      >
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

        <div class="select flex gap-2 mb-4">
          <div class="title-container">
            <div class="blue-box"></div>
            <span class="title-text">打卡位置</span>
          </div>
          <NormalMap ref="NormalMap"/>
        </div>
      </el-form>
    </div>
  </el-drawer>
</template>

<script>
import NormalMap from './map.vue'
import {selectCheckinInfoById} from '@/api/report/checkin.js'
export default {
  components: {
    NormalMap,
  },
  data() {
    return {
      visible: false, // 控制弹窗显示
      form: {},
      rules: {},
    };
  },
  methods: {
    formatCheckInDate(value) {
      if (!value) return "";
      if (Array.isArray(value) && value.length >= 3) {
        return `${value[0]}-${String(value[1]).padStart(2, "0")}-${String(
          value[2]
        ).padStart(2, "0")}`;
      }
      if (typeof value === "string") {
        return value.slice(0, 10);
      }
      return "";
    },
    showDrawer(statusId) {
      this.fetchCheckinPosition(statusId) 
      this.visible = true;
    },
    handleCancel() {
      this.visible = false;
    },

     // 根据id获取用户打卡位置信息
    async fetchCheckinPosition(statusId) {
      try {
        
        const response = await selectCheckinInfoById(statusId); 

        if (response.data.code === 1) {
          const data = response.data.data || {};
          const profile = data.profile || {};
          const user = data.user || {};
          const status = data.status || {};

          this.form = {
            userType: data.userType || "",
            name: data.name || user.name || "",
            age: data.age ?? profile.age ?? "",
            gender: data.gender || profile.gender || "",
            ethnicity: data.ethnicity || profile.ethnicity || "",
            department: data.department || profile.department || "",
            checkInDate: this.formatCheckInDate(
              data.checkInDate || status.checkInDate
            ),
          };

          const longitude = data.longitude ?? status.longitude;
          const latitude = data.latitude ?? status.latitude;
          const address =
            data.address ||
            [data.province || status.province, data.city || status.city, data.county || status.county, data.locationName || status.locationName]
              .filter(Boolean)
              .join("");
          this.$nextTick(() => {
            this.$refs.NormalMap?.getPosition(longitude, latitude, address);
          });
        } else {
          console.error("获取用户信息失败:", response.data.msg);
        }
      } catch (error) {
        console.error("请求出错:", error);
      }
    },
  },
};
</script>



<style  scoped>
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
.custom-drawer {
  height: 100%;
}
.form-container {
  padding: 20px;
  margin-top: 70px;
}
.title {
  position: fixed;
  top: 0;
  right: 0;
  width: 700px;
  background: #ffffff;
  padding-bottom: 25px;
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
