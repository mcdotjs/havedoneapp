<script lang="ts" setup>
import tailwindcss from "tailwindcss";
import { onMounted, reactive, ref, computed, watchEffect } from "vue";
import {
  Print,
  GetCount,
  StartStopwatch,
  StartTickerLoop,
  StopTicker,
} from "../wailsjs/go/main/App";

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
});

const count = ref(0);

const str = ref("");
const interval = ref("1s");
const timerDuration = ref("0s");
const lenInSeconds = ref(0);
const stop = () => {
  StopTicker();
};

function start() {
  StartTickerLoop(interval.value, timerDuration.value);
}

onMounted(() => {
  StartStopwatch("1s");

  setInterval(async () => {
    count.value = await GetCount();
    str.value = await Print();
  }, 100);
});

const colors = [
  { color: "#f56c6c", percentage: 20 },
  { color: "#e6a23c", percentage: 40 },
  { color: "#5cb87a", percentage: 60 },
  { color: "#1989fa", percentage: 80 },
  { color: "#6f7ad3", percentage: 100 },
];

const percentage = computed(() => (count.value / lenInSeconds.value) * 100);
watchEffect(() => {
  timerDuration.value = lenInSeconds.value + "s";
});
</script>
<template>
  <div class="flex flex-wrap justify-center items-center gap-16 py-12">
    <div class="text-5xl p-3 py-7">
      {{ str }}
    </div>
    <div class="">
      <el-progress
        type="dashboard"
        width="300"
        :percentage="percentage.toFixed()"
        :color="colors"
      />
    </div>
    <div class="slider-demo-block flex justify-center items-center w-full">
      <el-slider v-model="lenInSeconds" :min="10" :max="333" />
    </div>
    <main class="w-full h-screen">
      <div id="input" class="text-black">
        <!-- <div class="bg-gray-300 flex flex-row justify-center gap-8"> -->
        <!--   <div class="text-3xl"> -->
        <!--     Timer Duration -->
        <!--     <el-input -->
        <!--       id="name" -->
        <!--       v-model="timerDuration" -->
        <!--       autocomplete="off" -->
        <!--       class="text-2xl" -->
        <!--       type="text" -->
        <!--     /> -->
        <!---->
        <!--     <el-tag size="small">School</el-tag> -->
        <!--   </div> -->
        <!-- </div> -->
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
