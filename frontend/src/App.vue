<script lang="ts" setup>
import { onMounted, reactive, ref, computed, watchEffect } from "vue";
import { EventsOn } from "../wailsjs/runtime/runtime";
import {
  StartStopwatch,
  StartTickerLoop,
  StopTicker,
} from "../wailsjs/go/main/App";

const count = ref(0);
const str = ref("");
const remain = ref("");
const timerDuration = ref("30m");
const lenInSeconds = ref(1800);
const setFocusTime = (v: any) => {
  StopTicker();
  lenInSeconds.value = v;
  pauseModal.value = false;
};
const stop = () => {
  StopTicker();
};

function start() {
  pauseModal.value = false;
  StartTickerLoop("1s", timerDuration.value);
}

onMounted(() => {
  pauseModal.value = false;
  StartStopwatch("1s");

  EventsOn("timer-update", (data) => {
    str.value = data.HumanReadable;
    count.value = data.Count;
    remain.value = data.RemainingTime;
  });
});

const pausesInSeconds = [
  {
    asString: "10s",
    val: 10,
    type: "info",
    help: "Bugy bugy",
  },

  {
    asString: "10m",
    val: 600,
    type: "warning",

    help: "Little shajt",
  },

  {
    asString: "30m",
    val: 1800,
    type: "success",

    help: "Best chojse",
  },

  {
    asString: "1h",
    val: 3600,
    type: "danger",
    help: "What am I doing here?!",
  },
];

const colors = [
  { color: "#f56c6c", percentage: 20 },
  { color: "#e6a23c", percentage: 40 },
  { color: "#5cb87a", percentage: 60 },
  { color: "#1989fa", percentage: 80 },
  { color: "#6f7ad3", percentage: 100 },
];

const percentage = computed(
  () => ((count.value + 1) / lenInSeconds.value) * 100,
);
const pauseModal = ref(false);
const formattedRemain = computed(() => {
  const res = [];
  const arr = remain.value.split(".")[0].split("m");
  console.log("rrrrr", arr);
  if (arr.length == 1) {
    res[1] = arr[0];
    arr[0] = "0";
    arr[1] = res[1];
  }
  if (arr[1].startsWith("-")) {
    arr[1] = arr[1].substring(1);
  }
  if (arr[0].startsWith("-")) {
    arr[0] = arr[0].substring(1);
  }

  if (arr[1].length == 1) {
    arr[1] = "0" + arr[1];
    console.log("11111", arr[1]);
  }

  if (arr[0].length == 1) {
    arr[0] = "0" + arr[0];
    console.log("0000", arr[0]);
  }
  return `${arr[0]}:${arr[1]}`;
});

watchEffect(() => {
  timerDuration.value = lenInSeconds.value + "s";
  if (count.value == lenInSeconds.value + 1) {
    pauseModal.value = true;
  }
});
</script>
<template>
  <div
    v-if="pauseModal"
    class="text-7xl absolute h-48 w-96 top-0 z-20 right-0 bg-red-500 rounded-lg m-12 flex justify-center items-center"
  >
    Pauza!!
    <span
      class="absolute top-2 right-3 text-xs cursor-pointer"
      @click="pauseModal = false"
      >ouyeeeh</span
    >
  </div>
  <div class="flex flex-col justify-center items-center gap-16 py-12">
    <div class="text-5xl p-3 py-7">
      {{ str }}
    </div>
    <div
      class="slider-demo-block flex flex-col justify-center items-center w-full text-7xl"
    >
      <div class="text-3xl italic pb-3">Fockus Time!!</div>
      <div class="flex justify-center items-center gap-3">
        <el-button
          v-for="item in pausesInSeconds"
          :key="item.val"
          :label="item.asString"
          @click="setFocusTime(item.val)"
          :type="item.type"
          size="large"
          link
          class="text-5xl relative"
        >
          {{ item.asString
          }}<span class="text-xs absolute -bottom-3 left-2 w-5">{{
            item.help
          }}</span></el-button
        >
      </div>
    </div>
    <main class="w-full flex flex-col justify-center items-center">
      <div class="relative">
        <el-progress
          type="dashboard"
          :width="300"
          :show-text="false"
          :percentage="Number(percentage.toFixed())"
          :color="colors"
        />
        <div
          class="absolute top-0 h-full w-full flex justify-center items-center text-4xl"
        >
          {{ formattedRemain }}
        </div>
      </div>
      <div id="input" class="text-black">
        <el-button
          type="success"
          size="large"
          link
          class="text-5xl"
          @click="start"
          >START</el-button
        >

        <el-button
          type="warning"
          size="large"
          link
          class="text-5xl"
          @click="stop"
          >STOP</el-button
        >
      </div>
    </main>
  </div>
</template>

<style>
.slider-demo-block {
  max-width: 900px;
  display: flex;
  align-items: center;
}
.slider-demo-block .el-slider {
  margin-top: 0;
  margin-left: 12px;
}
.slider-demo-block .demonstration {
  font-size: 14px;
  color: var(--el-text-color-secondary);
  line-height: 44px;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-bottom: 0;
}
.slider-demo-block .demonstration + .el-slider {
  flex: 0 0 70%;
}
</style>
