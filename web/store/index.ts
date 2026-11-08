// import Vue from 'vue'
// import Vuex from 'vuex'
// Vue.use(Vuex);
// const store = new Vuex.Store({
// 	state: {
// 		list: [{
// 				pagePath: "/pages/index/index",
// 				iconPath: "home",
// 				selectedIconPath: "home-fill",
// 				text: '首页',
// 				count: 2,
// 				isDot: true,
// 				customIcon: false,
// 			},
// 			{
// 				pagePath: "/pages/save/save",
// 				iconPath: "photo",
// 				selectedIconPath: "photo-fill",
// 				text: '省钱',
// 				customIcon: false,
// 			},
// 			{
// 				pagePath: "/pages/scan/scan",
// 				iconPath: "/static/scan.png",
// 				selectedIconPath: "/static/scan-selected.png",
// 				text: '扫码',
// 				midButton: true,
// 				customIcon: false,
// 			},
// 			{
// 				pagePath: "/pages/money/money",
// 				iconPath: "play-right",
// 				selectedIconPath: "play-right-fill",
// 				text: '钱包',
// 				customIcon: false,
// 			},
// 			{
// 				pagePath: "/pages/user/user",
// 				iconPath: "account",
// 				selectedIconPath: "account-fill",
// 				text: '个人中心',
// 				count: 23,
// 				isDot: false,
// 				customIcon: false,
// 			},
// 		],
// 	}
// })

// export default store
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

// 定义状态类型
interface State {
  jwt: string | null;
  userName: string;
  clockFormParams: {
    city: string;
    county: string;
    health: boolean;
    latitude: number;
    longitude: number;
    locationName: string;
    province: string;
  };
  diseasescoring: Diseasescoring[];
  scoringHybrid: ScoringHybrid | null;
}
interface Diseasescoring {
  diseaseTypeId: number;
  diseaseTypeName: string;
  score: number;
  scoringId: number;
  statusId: number;
  symptomTime: number[];
  reason?: string;
  evidence?: string[];
  expertScore?: number;
}
interface ScoringHybrid {
  isAI?: boolean;
  disagreement?: boolean;
  needsHumanConfirm?: boolean;
  confirmed?: boolean;
  llmDegraded?: boolean;
  expertTop?: string[];
  aiTop?: string[];
  message?: string;
}

// 定义 mutations 类型
interface Mutations {
  setJwt(state: State, jwt: string): void;
  setUserName(state: State, userName: string): void;
  setClockFormParams(state: State, clockFormParams: State['clockFormParams']): void;
  setDiseasescoring(state: State, diseascoring: Diseasescoring[]): void;
  setScoringHybrid(state: State, hybrid: ScoringHybrid | null): void;
}

// 定义 actions 类型
interface Actions {
  setJwt({ commit }: { commit: (type: keyof Mutations, payload: string) => void }, jwt: string): void;
  setUserName({ commit }: { commit: (type: keyof Mutations, payload: string) => void }, userName: string): void;
  setClockFormParams({ commit }: { commit: (type: keyof Mutations, payload: State['clockFormParams']) => void }, clockFormParams: State['clockFormParams']): void;
  setDiseasescoring({ commit }: { commit: (type: keyof Mutations, payload: { diseascoring: Diseasescoring[] }) => void }, diseascoring: Diseasescoring[]): void;
  setScoringHybrid({ commit }: { commit: (type: keyof Mutations, payload: ScoringHybrid | null) => void }, hybrid: ScoringHybrid | null): void;
}

function readStoredUserName(): string {
  try {
    return uni.getStorageSync('userName') || ''
  } catch {
    return ''
  }
}

// 创建 store
const store = new Vuex.Store<State>({
  state: {
    jwt: null, // 存储 JWT 信息
    userName: readStoredUserName(),
    clockFormParams: {
      city: '',
      county: '',
      health: false,
      latitude: 0,
      longitude: 0,
      locationName: '',
      province: ''
    },
    diseasescoring: [],
    scoringHybrid: null,
  },
  mutations: {
    setJwt(state, jwt: string) {
      state.jwt = jwt;
    },
    setUserName(state, userName: string) {
      state.userName = userName || '';
      try {
        if (userName) {
          uni.setStorageSync('userName', userName);
        } else {
          uni.removeStorageSync('userName');
        }
      } catch {
        // ignore storage failures in mini-program tooling
      }
    },
    setClockFormParams(state, clockFormParams: State['clockFormParams']) {
      state.clockFormParams = { ...clockFormParams };
    },
    setDiseasescoring(state, diseascoring: Diseasescoring[]) {
      state.diseasescoring = diseascoring;
    },
    setScoringHybrid(state, hybrid: ScoringHybrid | null) {
      state.scoringHybrid = hybrid;
    },
  },
  actions: {
    setJwt({ commit }, jwt: string) {
      commit('setJwt', jwt);
    },
    setUserName({ commit }, userName: string) {
      commit('setUserName', userName);
    },
    setClockFormParams({ commit }, clockFormParams: State['clockFormParams']) {
      commit('setClockFormParams', clockFormParams);
    },
    setDiseasescoring({ commit }, diseascoring: Diseasescoring[]) {
      commit('setDiseasescoring', diseascoring);
    },
    setScoringHybrid({ commit }, hybrid: ScoringHybrid | null) {
      commit('setScoringHybrid', hybrid);
    },
  },
});

export default store;



