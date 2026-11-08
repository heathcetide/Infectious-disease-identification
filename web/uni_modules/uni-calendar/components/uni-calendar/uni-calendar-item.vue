<template>
	<view
		class="uni-calendar-item__weeks-box"
		:class="{
			'is-disable': weeks.disable,
			'is-today': weeks.isDay,
			'is-selected': isSelected,
		}"
		@click="choiceDate(weeks)"
	>
		<view class="day-cell">
			<!-- conic status ring: health / risk / diagnose / examine -->
			<view
				v-if="ringClass"
				class="day-ring"
				:class="[ringClass, { 'ring-lift': isSelected || weeks.isDay }]"
			/>
			<view
				v-if="weeks.isDay && !isSelected"
				class="day-today-ring"
			/>
			<text
				class="day-num"
				:class="{
					'num-mute': weeks.disable,
					'num-today': weeks.isDay,
					'num-selected': isSelected,
					'num-marked': !!ringClass,
				}"
			>{{ weeks.date }}</text>
		</view>
	</view>
</template>

<script>
	export default {
		emits: ['change'],
		props: {
			weeks: {
				type: Object,
				default() {
					return {}
				},
			},
			calendar: {
				type: Object,
				default: () => ({}),
			},
			selected: {
				type: Array,
				default: () => [],
			},
			lunar: {
				type: Boolean,
				default: false,
			},
		},
		computed: {
			isSelected() {
				return this.calendar && this.calendar.fullDate === this.weeks.fullDate
			},
			statusInfo() {
				const info = this.weeks && this.weeks.extraInfo && this.weeks.extraInfo.info
				return Array.isArray(info) && info.length >= 3 ? info : null
			},
			/**
			 * info = [health, diagnose, examine]
			 * 健康 → 全绿
			 * 不适 → 全红
			 * 不适+诊断 → 红/橙 各半
			 * 再加检测 → 紫/红/橙 三等分
			 * 健康+诊断/检测 同样按段数拆开，底色用绿
			 */
			ringClass() {
				const info = this.statusInfo
				if (!info) return ''
				const healthy = Number(info[0]) === 1
				const diagnose = Number(info[1]) === 1
				const examine = Number(info[2]) === 1
				const flags = (diagnose ? 1 : 0) + (examine ? 1 : 0)

				if (healthy) {
					if (flags === 0) return 'ring-health'
					if (diagnose && examine) return 'ring-health-diagnose-examine'
					if (diagnose) return 'ring-health-diagnose'
					return 'ring-health-examine'
				}
				// 不适
				if (flags === 0) return 'ring-risk'
				if (diagnose && examine) return 'ring-risk-diagnose-examine'
				if (diagnose) return 'ring-risk-diagnose'
				return 'ring-risk-examine'
			},
		},
		methods: {
			choiceDate(weeks) {
				this.$emit('change', weeks)
			},
		},
	}
</script>

<style lang="scss" scoped>
	$c-health: #22c55e;
	$c-risk: #ef4444;
	$c-diagnose: #f97316;
	$c-examine: #8b5cf6;
	$c-primary: #285ac8;

	.uni-calendar-item__weeks-box {
		flex: 1;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: column;
		justify-content: center;
		align-items: center;
		background: transparent !important;
	}

	.day-cell {
		position: relative;
		width: 78rpx;
		height: 78rpx;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		justify-content: center;
		align-items: center;
	}

	.day-ring {
		position: absolute;
		left: 50%;
		top: 50%;
		width: 64rpx;
		height: 64rpx;
		margin-left: -32rpx;
		margin-top: -32rpx;
		border-radius: 50%;
		z-index: 1;
		transition: transform 0.2s ease;
	}

	.day-ring::before {
		content: '';
		position: absolute;
		left: 50%;
		top: 50%;
		width: 72%;
		height: 72%;
		transform: translate(-50%, -50%);
		border-radius: 50%;
		background: #ffffff;
		z-index: 1;
	}

	.ring-lift {
		transform: scale(1.05);
	}

	.ring-health {
		background: $c-health;
	}
	.ring-risk {
		background: $c-risk;
	}
	.ring-health-diagnose {
		background: conic-gradient(from -90deg, $c-diagnose 0 180deg, $c-health 180deg 360deg);
	}
	.ring-health-examine {
		background: conic-gradient(from -90deg, $c-examine 0 180deg, $c-health 180deg 360deg);
	}
	.ring-health-diagnose-examine {
		background: conic-gradient(
			from -90deg,
			$c-examine 0 120deg,
			$c-health 120deg 240deg,
			$c-diagnose 240deg 360deg
		);
	}
	.ring-risk-diagnose {
		background: conic-gradient(from -90deg, $c-diagnose 0 180deg, $c-risk 180deg 360deg);
	}
	.ring-risk-examine {
		background: conic-gradient(from -90deg, $c-examine 0 180deg, $c-risk 180deg 360deg);
	}
	.ring-risk-diagnose-examine {
		background: conic-gradient(
			from -90deg,
			$c-examine 0 120deg,
			$c-risk 120deg 240deg,
			$c-diagnose 240deg 360deg
		);
	}

	.day-today-ring {
		position: absolute;
		left: 50%;
		top: 50%;
		width: 44rpx;
		height: 44rpx;
		margin-left: -22rpx;
		margin-top: -22rpx;
		border-radius: 50%;
		z-index: 2;
		box-shadow: inset 0 0 0 2rpx rgba(40, 90, 200, 0.55);
		background: transparent;
	}

	.day-num {
		position: relative;
		z-index: 3;
		font-size: 28rpx;
		color: #1f2937;
		line-height: 1;
		font-weight: 500;
	}

	.num-mute {
		color: #c0c4cc;
		font-weight: 400;
	}

	.num-marked {
		font-weight: 600;
	}

	.num-today {
		color: $c-primary;
		font-weight: 700;
	}

	.num-selected {
		color: $c-primary !important;
		font-weight: 700;
	}

	.is-disable .day-ring,
	.is-disable .day-today-ring {
		opacity: 0.35;
	}

	.is-today,
	.is-selected {
		background-color: transparent !important;
		opacity: 1 !important;
	}
</style>
