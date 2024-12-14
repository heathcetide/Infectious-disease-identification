<script>
import { reportSubmitService } from '../../services/experiment'
import { getLocalReportService } from '../../services/clock'
export default {
  data() {
    return {
      formData: {
        date: '',
        stoolReport: [],   // 粪便检查报告文件数组
        vomitReport: [],   // 呕吐物检查报告文件数组
        bloodReport: []    // 血液检查报告文件数组
      },
      reportForm: {
        isSerologicalTestDone: false,
        isVirusAntigenTestDone: false,
        isVirusNucleicAcidTestDone: false,
        isVirusCultureIsolationDone: false,
		pathogenicTestResults: '',
        labTestFiles: []  // 新增的 labTestFiles 数组
      },
	  isViewMode: false,
    };
  },
  onLoad(options) {
    // `options` 是页面加载时获取的传参对象
    const { date } = options;
  	console.log('dedao',date,'那么isViewMode=',this.isViewMode)
    if (date) {
      this.isViewMode = true;
      // 调用获取服务方法以查询数据
      this.getDiagnosisData(date);
    }
  },
  methods: {
	async getDiagnosisData(date) {
		getLocalReportService(date)
				.then(res => {
					console.log(res.data)
					this.reportForm = res.data
				    this.reportForm.labTestFiles.forEach(file => {
				        const field = file.fileId; // 假设fileId是你想用的字段
				        const newFile = {
				            name: `${file.specimenType}_report_${field}.pdf`, // 文件名可自定义
				            extname: "pdf", // 假设都是pdf格式，若有不同可根据实际情况修改
				            url: `https://ds.sccdc.cn/common/file/getLabTestFile/${field}`,
				        };
				
				        if (file.specimenType === "粪便") {
				            this.formData.stoolReport.push(newFile);
				        } else if (file.specimenType === "呕吐物") {
				            this.formData.vomitReport.push(newFile);
				        } else if (file.specimenType === "血液") {
				            this.formData.bloodReport.push(newFile);
				        }
				    });
					console.log(this.formData.stoolReport)
				})
				.catch(error => this.submitError(error));
	},
    handleFileChange(event, reportType) {
		const selectFile =  [{
			path: event.tempFilePaths[0],
			name: event.tempFiles[0].name,
		}]
        this.formData[reportType] = [...this.formData[reportType], ...selectFile];
    
        console.log(this.formData);
    },
    onDateChange(e) {
      this.formData.date = e.detail.value;
    },
    chooseFile(reportType) {
      wx.chooseMessageFile({
        count: 3, // 最多可以选择3个文件
        type: 'file',
        extension: ['pdf', 'jpg', 'png'], // 支持的文件类型
        success: (res) => {
          const files = res.tempFiles.map(file => ({
            path: file.path,
            name: file.name
          }));
          this.formData[reportType] = [...this.formData[reportType], ...files];
        }
      });
    },

    uploadFile(filePath, reportType, index) {
      return new Promise((resolve, reject) => {
        wx.uploadFile({
          url: 'https://ds.sccdc.cn/railwayemployee/labtest/file',
          filePath: filePath,
          name: 'file',
          formData: {
            labTestType: reportType,
            index: index
          },
          success: (res) => {
            console.log(`文件上传成功: ${res.data}`);
			console.log(JSON.parse(res.data).data)
            const fileId = parseInt(JSON.parse(res.data).data);
            this.reportForm.labTestFiles.push(fileId);
			console.log('添加了id',this.reportForm.labTestFiles,"666",fileId)
            resolve(res);
          },
          fail: (err) => {
            console.error(`文件上传失败: ${err}`);
            reject(err);
          }
        });
      });
    },

    handleSubmit() {
      wx.showModal({
        content: '提交后不可修改，是否确认提交？',
        cancelText: '取消',
        confirmText: '确认提交',
        success: async (res) => {
          if (res.confirm) {
            // 上传所有报告的文件
            const uploadPromises = [];
            // 上传粪便检查报告
            this.formData.stoolReport.forEach((file, index) => {
              uploadPromises.push(this.uploadFile(file.path, '粪便', index));
            });
            // 上传呕吐物检查报告
            this.formData.vomitReport.forEach((file, index) => {
              uploadPromises.push(this.uploadFile(file.path, '呕吐物', index));
            });

            // 上传血液检查报告
            this.formData.bloodReport.forEach((file, index) => {
              uploadPromises.push(this.uploadFile(file.path, '血', index));
            });
            try {
              // 等待所有文件上传完成
              await Promise.all(uploadPromises);
              console.log('所有文件上传完成',this.reportForm);
				
              // 提交表单
              reportSubmitService(this.reportForm)
                .then((res) => this.submitSuccess(res))
                .catch((error) => this.submitError(error));
				
              wx.showToast({
                title: '提交成功',
                icon: 'success',
                duration: 2000
              });
            } catch (error) {
              console.error('上传过程出现错误: ', error);
              wx.showToast({
                title: error,
				content: error,
                icon: 'none',
                duration: 2000
              });
            }
          } else if (res.cancel) {
            console.log('用户点击取消');
          }
        }
      });
    },
    submitError(error) {
      uni.showToast({
        title: '信息提交失败',
        icon: 'none',
        duration: 1500
      });
    },
    submitSuccess(response) {
	  console.log("注意注意",response)
      if (response.code == 1) {
        uni.showToast({
          title: '信息提交成功',
          icon: 'success',
          duration: 1800
        });
      } else {
        uni.showModal({
          title: '信息提交失败',
          content: response.msg,
          showCancel: false,
          confirmText: '确定',
          success: function (res) {}
        });
      }
    },
    handleReport(propName) {
      console.log(`Changing ${propName}:`, this.reportForm[propName]);
      this.reportForm[propName] = !this.reportForm[propName];
    },
    isImage(filePath) {
      const extension = filePath.split('.').pop().toLowerCase();
      return ['jpg', 'png'].includes(extension);
    }
  }
}
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="content">
      <!-- 病原学及血清学检查 -->
      <view class="section">
        <view class="sections1">
          <view class="section-title">
            <text class="label">病原学及血清学检查（可多选）</text>
          </view>
          <view class="checkboxs">
            <checkbox-group>
              <view class="order-options">
                <label class="checkbox-item" @tap="handleReport('isVirusAntigenTestDone')">
                  <checkbox color="#285AC8" :checked="reportForm.isVirusAntigenTestDone" :disabled="isViewMode"/>
                  病毒抗原检测
                </label>
                <label class="checkbox-item" @tap="handleReport('isVirusNucleicAcidTestDone')" >
                  <checkbox color="#285AC8" :checked="reportForm.isVirusNucleicAcidTestDone" :disabled="isViewMode"/>
                  病毒核酸检测
                </label>
              </view>
              <view class="order-options">
                <label class="checkbox-item" @tap="handleReport('isVirusCultureIsolationDone')">
                  <checkbox color="#285AC8" :checked="reportForm.isVirusCultureIsolationDone" :disabled="isViewMode"/>
                  病毒培养分离
                </label>
                <label class="checkbox-item" @tap="handleReport('isSerologicalTestDone')">
                  <checkbox color="#285AC8" :checked="reportForm.isSerologicalTestDone" :disabled="isViewMode"/>
                  血清学检测
                </label>
              </view>
            </checkbox-group>
          </view>

          <view class="row" v-show="!isViewMode">
            <view >
              <view class="section-title">
                <text class="label">提交时间</text>
              </view>
              <picker mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange" class="time-picker">
                <view class="picker">
                  {{ formData.date || '请选择日期' }}
                </view>
              </picker>
            </view>
          </view>
		  <view class="row" v-show="!isViewMode">
			  <view>
				<view class="section-title">
				    <text class="label">病原学及血清学检查结果</text>
					<input type="text" placeholder="请输入检查结果" v-model="reportForm.pathogenicTestResults" class="text-input" style="margin-left:6px; margin-top:25px;float: right; width: 200px;"/>
				</view>
			  </view>
		  </view>
		  
		  <!-- 粪便检查报告 -->
		  <view class="row">
		    <view class="form-item">
		      <view class="section-title">
		        <text class="label">粪便检查报告：</text>
		      </view>
            <view class="example-body">
                <!-- 粪便检查报告 -->
                <uni-file-picker 
     				:disabled="isViewMode"
				    :value="formData.stoolReport"
                    :limit="3" 
					:show-type="'grid'"
                    :file-type="['image', 'pdf']"
                    :filesList="formData.stoolReport" 
                    fileMediatype="image" 
                    @select="handleFileChange($event, 'stoolReport')"
                    title="最多选择3个文件 (PDF, JPG, PNG)">
                </uni-file-picker>
            </view>
		    </view>
		  </view>
          <!-- 呕吐物检查报告 -->
          <view class="row">
            <view class="form-item">
              <view class="section-title">
                <text class="label">呕吐物检查报告：</text>
              </view>
			  <view class="example-body">
			      <!-- 呕吐物检查报告 -->
			      <uni-file-picker 
			          :disabled="isViewMode"
			          :value="formData.vomitReport"
					  :show-type="'grid'"
					  fileMediatype="image" 
			          :limit="3" 
			          :file-type="['image', 'pdf']" 
			          :filesList="formData.vomitReport" 
			          @select="handleFileChange($event, 'vomitReport')"
			          title="最多选择3个文件 (PDF, JPG, PNG)">
			      </uni-file-picker>
			  </view>
            </view>
          </view>

          <!-- 血液检查报告 -->
          <view class="row">
            <view class="form-item">
              <view class="section-title">
                <text class="label">血液检查报告：</text>
              </view>
			  <view class="example-body">
			      <!-- 血液检查报告 -->
			      <uni-file-picker 
			          :disabled="isViewMode"
			          :value="formData.bloodReport"
					  fileMediatype="image" 
					  :show-type="'grid'"
			          :limit="3" 
			          :file-type="['image', 'pdf']" 
			          :files-list="formData.bloodReport" 
			          @select="handleFileChange($event, 'bloodReport')"
			          title="最多选择3个文件 (PDF, JPG, PNG)">
			      </uni-file-picker>
			  </view>
            </view>
          </view>

        </view>
        <view class="button-container">
          <button class="single-submit-button" @tap="handleSubmit()">提交</button>
        </view>
      </view>
    </view>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
</template>

<style>
.container {
	width: 100%;
	height: 100%;
	position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
	
.background-image {
	position: absolute; /* 设置为绝对定位，使其占据整个容器 */
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	width: 100%;
	height: 100%;
	z-index: -1; /* 将图片置于内容下方 */
}
.tabbar{
	z-index: 999;
}
.content{
	position: relative;
	padding-top: 14vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 13vh;
}
.header {
  position: relative;
  text-align: left;
  margin-bottom: 20px;
  font-size: 26px;
  font-weight: bold;
  padding: 0 10px;
}
.form-item4{
	width: 100%;
	margin: 3px 2px;
	opacity: 1;
	border-radius: 5px;
	background: rgba(255, 255, 255, 1);
	padding: 8px;
	margin-bottom: 8px;
}
.title {
  position: relative;
}
  .title::after {
    content: "";
    position: absolute;
    bottom: -4px;
    left: 0;
    width: 100%;
    height: 6px;
    background-color: #285AC8;
  }
.section {
  margin-bottom: 20px;
  padding: 12px;
  border-radius: 8px;
  color: #333333;
}
.section-title{
	margin-left: -3px;
	padding-left: 4px;
	border-left: 8px solid #285AC8;
}
.row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 24px;
  margin-bottom: 36px;
}
.checkboxs{
	margin: 7px 30px;
}
.order-options{
	display: block; 
	margin: 15px 0; 
	display: flex; 
	justify-content: space-between; 
}
.inner-checkboxs{
	margin-left: 4vh;
}
.label {
  font-size: 16px;
  color: #333333;
  display: block;
}

.options {
  display: flex;
}
.time-picker{
	margin-top: -18px;
	margin-left: 180px;
}
.single-submit-button{
	height: 44px;
	background-color: #285AC8;
	color: #ffffff;
	text-align: center;
	border-radius: 8px;
	font-size: 16px;
	margin: 0 24px;
}
.center-button {
  width: 90px;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 14px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.example-body {
	padding: 10px;
	padding-top: 0;
}

.custom-image-box {
	/* #ifndef APP-NVUE */
	display: flex;
	/* #endif */
	flex-direction: row;
	justify-content: space-between;
	align-items: center;
}

.text {
	font-size: 14px;
	color: #333;
}
input{
	opacity: 0.7;
}
</style>