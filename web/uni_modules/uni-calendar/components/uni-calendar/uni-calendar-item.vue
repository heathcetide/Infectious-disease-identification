<template>
	<view class="uni-calendar-item__weeks-box" :class="{
		'uni-calendar-item--disable':weeks.disable,
		'uni-calendar-item--isDay':calendar.fullDate === weeks.fullDate && weeks.isDay,
		'uni-calendar-item--checked':(calendar.fullDate === weeks.fullDate && !weeks.isDay) ,
		'uni-calendar-item--before-checked':weeks.beforeMultiple,
		'uni-calendar-item--multiple': weeks.multiple,
		'uni-calendar-item--after-checked':weeks.afterMultiple,
		}"
	 @click="choiceDate(weeks)">
		<view class="uni-calendar-item__weeks-box-item">
			<text v-if="selected&&weeks.extraInfo" class="uni-calendar-item__weeks-box-circle"></text>
			<text class="uni-calendar-item__weeks-box-text" :class="{
				'uni-calendar-item--isDay-text': weeks.isDay,
				'uni-calendar-item--isDay':calendar.fullDate === weeks.fullDate && weeks.isDay,
				'uni-calendar-item--checked':calendar.fullDate === weeks.fullDate && !weeks.isDay,
				'uni-calendar-item--before-checked':weeks.beforeMultiple,
				'uni-calendar-item--multiple': weeks.multiple,
				'uni-calendar-item--after-checked':weeks.afterMultiple,
				'uni-calendar-item--disable':weeks.disable,
				}">{{weeks.date}}</text>
			<text v-if="!lunar&&!weeks.extraInfo && weeks.isDay" class="uni-calendar-item__weeks-lunar-text" :class="{
				'uni-calendar-item--isDay-text':weeks.isDay,
				'uni-calendar-item--isDay':calendar.fullDate === weeks.fullDate && weeks.isDay,
				'uni-calendar-item--checked':calendar.fullDate === weeks.fullDate && !weeks.isDay,
				'uni-calendar-item--before-checked':weeks.beforeMultiple,
				'uni-calendar-item--multiple': weeks.multiple,
				'uni-calendar-item--after-checked':weeks.afterMultiple,
				}">{{todayText}}</text>
			<text v-if="lunar&&!weeks.extraInfo" class="uni-calendar-item__weeks-lunar-text" :class="{
				'uni-calendar-item--isDay-text':weeks.isDay,
				'uni-calendar-item--isDay':calendar.fullDate === weeks.fullDate && weeks.isDay,
				'uni-calendar-item--checked':calendar.fullDate === weeks.fullDate && !weeks.isDay,
				'uni-calendar-item--before-checked':weeks.beforeMultiple,
				'uni-calendar-item--multiple': weeks.multiple,
				'uni-calendar-item--after-checked':weeks.afterMultiple,
				'uni-calendar-item--disable':weeks.disable,
				}">{{weeks.isDay ? todayText : (weeks.lunar.IDayCn === '初一'?weeks.lunar.IMonthCn:weeks.lunar.IDayCn)}}</text>
			<text v-if="weeks.extraInfo&&weeks.extraInfo.info" class="uni-calendar-item__weeks-lunar-text" :class="{
				'uni-calendar-item--extra':weeks.extraInfo.info,
				'uni-calendar-item--isDay-text':weeks.isDay,
				'uni-calendar-item--isDay':calendar.fullDate === weeks.fullDate && weeks.isDay,
				'uni-calendar-item--checked':calendar.fullDate === weeks.fullDate && !weeks.isDay,
				'uni-calendar-item--before-checked':weeks.beforeMultiple,
				'uni-calendar-item--multiple': weeks.multiple,
				'uni-calendar-item--after-checked':weeks.afterMultiple,
				'uni-calendar-item--disable':weeks.disable,
				'has-unhealth-examine-diagnose-status':shouldUnhealthExamineDiagnoseStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-health-status': shouldHealthStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-unhealth-diagnose-status': shouldUnhealthDiagnoseStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-unhealth-status': shouldUnhealthStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-unhealth-examine-status': shouldUnhealthExamineStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-health-diagnose-status': shouldHealthDiagnoseStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-health-examine-status': shouldHealthExamineStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				'has-health-examine-diagnose-status': shouldHealthExamineDiagnoseStatusBorder(weeks.extraInfo&&weeks.extraInfo.info&&weeks.fullDate),
				}">
				<!-- {{weeks.extraInfo.info}} -->
<!-- 				<text style="display:inline-block; width:20rpx; height:20rpx;border-radius: 20rpx;" v-for="(item, index) in weeks.extraInfo.info"
				:style="{background:returnColor(item)}"></text>
				<view style="display:inline-block; width:20rpx; height:20rpx; background-color: aqua;"></view> -->
				</text>
		</view>
	</view>
</template>

<script>
	import { initVueI18n } from '@dcloudio/uni-i18n'
	import i18nMessages from './i18n/index.js'
	const {	t	} = initVueI18n(i18nMessages)

	export default {
		emits:['change'],
		props: {
			weeks: {
				type: Object,
				default () {
					return {}
				}
			},
			calendar: {
				type: Object,
				default: () => {
					return {}
				}
			},
			selected: {
				type: Array,
				default: () => {
					return []
				}
			},
			lunar: {
				type: Boolean,
				default: false
			}
		},
		computed: {
			todayText() {
				return t("uni-calender.today")
			},
		},
		methods: {
			shouldHealthExamineDiagnoseStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==1&&this.weeks.extraInfo.info[1]==1&&this.weeks.extraInfo.info[2]==1){
					return weeks;
				}
			    return false;
			},
			shouldHealthExamineStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==1&&this.weeks.extraInfo.info[1]==0&&this.weeks.extraInfo.info[2]==1){
					return weeks;
				}
			    return false;
			},
			shouldHealthDiagnoseStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==1&&this.weeks.extraInfo.info[1]==1&&this.weeks.extraInfo.info[2]==0){
					return weeks;
				}
			    return false;
			},
			shouldUnhealthExamineStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==0&&this.weeks.extraInfo.info[1]==0&&this.weeks.extraInfo.info[2]==1){
					return weeks;
				}
			    return false;
			},
			shouldUnhealthStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==0&&this.weeks.extraInfo.info[1]==0&&this.weeks.extraInfo.info[2]==0){
					return weeks;
				}
			    return false;
			},
			shouldUnhealthDiagnoseStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==0&&this.weeks.extraInfo.info[1]==1&&this.weeks.extraInfo.info[2]==0){
					return weeks;
				}
			    return false;
			},
			shouldHealthStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==1&&this.weeks.extraInfo.info[1]==0&&this.weeks.extraInfo.info[2]==0){
					return weeks;
				}
			    return false;
			},
			
			shouldUnhealthExamineDiagnoseStatusBorder(weeks) {
				if(this.weeks.extraInfo.info[0]==0&&this.weeks.extraInfo.info[1]==1&&this.weeks.extraInfo.info[2]==1){
					return weeks;
				}
			    return false;
			},
			choiceDate(weeks) {
				this.$emit('change', weeks)
			},
			returnColor(type){
				let typeColor = null
				switch (type) {
					case 1:
						typeColor = '#1677FF';
						break;
					case 2:
						typeColor = '#57BE2D';
						break;
					case 3:
						typeColor = '#FA6400';
						break;
					case 4:
						typeColor = '#F7B500';
						break;
					case 5:
						typeColor = '#E02020';
						break;
					case 6:
						typeColor = '#53D1F0';
						break;
					case 7:
					case 8:
						typeColor = '#1677FF';
						break;
					default:
						typeColor = '#000000'; // 默认颜色
						break;
				}
				return typeColor
			}
		}
	}
</script>

<style lang="scss" scoped>
	$uni-font-size-base:14px;
	$uni-text-color:#333;
	$uni-font-size-sm:12px;
	$uni-color-error: #e43d33;
	$uni-opacity-disabled: 0.3;
	$uni-text-color-disable:#c0c0c0;
	$uni-primary: #2979ff !default;
	.uni-calendar-item__weeks-box {
		flex: 1;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.uni-calendar-item__weeks-box-text {
		font-size: $uni-font-size-base;
		color: $uni-text-color;
	}

	.uni-calendar-item__weeks-lunar-text {
		font-size: $uni-font-size-sm;
		color: $uni-text-color;
	}

	.uni-calendar-item__weeks-box-item {
		position: relative;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 100rpx;
		height: 100rpx;
	}

	.uni-calendar-item__weeks-box-circle {
		position: absolute;
		top: 5px;
		right: 5px;
		width: 8px;
		height: 8px;
		border-radius: 8px;
		// background-color: $uni-color-error;
	}

	.uni-calendar-item--disable {
		background-color: rgba(249, 249, 249, $uni-opacity-disabled);
		color: $uni-text-color-disable;
	}

	.uni-calendar-item--isDay-text {
		color: $uni-primary;
	}

	.uni-calendar-item--isDay {
		background-color: $uni-primary;
		opacity: 0.8;
		color: #fff;
	}

	.uni-calendar-item--extra {
		color: $uni-color-error;
		opacity: 0.8;
	}

	.uni-calendar-item--checked {
		background-color: $uni-primary;
		color: #fff;
		opacity: 0.8;
	}

	.uni-calendar-item--multiple {
		background-color: $uni-primary;
		color: #fff;
		opacity: 0.8;
	}
	.uni-calendar-item--before-checked {
		background-color: #ff5a5f;
		color: #fff;
	}
	.uni-calendar-item--after-checked {
		background-color: #ff5a5f;
		color: #fff;
	}
.has-red-border {
	position: absolute;
    width: 60%;
    height: 60%;
    border-radius: 50%;
    background: conic-gradient(
        red 0deg 120deg,    /* 左上角红色 */
        blue 120deg 240deg, /* 右上角蓝色 */
        green 240deg 360deg /* 底部绿色 */
    );
    display: flex;
    align-items: center;
    justify-content: center;
}
.has-red-border::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}
.has-health-status {
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    green 0deg 120deg,    /* 左上角红色 */
	    green 120deg 240deg, /* 右上角蓝色 */
	    green 240deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-health-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}
.has-unhealth-status{
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    yellow 0deg 120deg,    /* 左上角红色 */
	    yellow 120deg 240deg, /* 右上角蓝色 */
	    yellow 240deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-unhealth-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}
.has-unhealth-diagnose-status{
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    red 0deg 180deg,    /* 左上角红色 */
	    yellow 180deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-unhealth-diagnose-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}
.has-unhealth-examine-status{
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    purple 0deg 180deg,    /* 左上角红色 */
	    yellow 180deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-unhealth-examine-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}

.has-unhealth-examine-diagnose-status {
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    purple 0deg 120deg,    /* 左上角红色 */
	    yellow 120deg 240deg, /* 右上角蓝色 */
	    red 240deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-unhealth-examine-diagnose-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}

.has-health-diagnose-status{
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    red 0deg 180deg,    /* 左上角红色 */
	    white 180deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-health-diagnose-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}

.has-health-examine-status{
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    purple 0deg 180deg,    /* 左上角红色 */
	    white 180deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-health-examine-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}

.has-health-examine-diagnose-status {
	position: absolute;
	width: 60%;
	height: 60%;
	border-radius: 50%;
	background: conic-gradient(
	    purple 0deg 120deg,    /* 左上角红色 */
	    white 120deg 240deg, /* 右上角蓝色 */
	    red 240deg 360deg /* 底部绿色 */
	);
	display: flex;
	align-items: center;
	justify-content: center;
}
.has-health-examine-diagnose-status::before{
	content: '';
	position: absolute;
            width: 80%;  /* 调小内圆的大小 */
            height: 80%; /* 调小内圆的大小 */
	background-color: white; /* 圆形内部颜色 */
	border-radius: 50%;
	z-index: 1;
}
</style>
