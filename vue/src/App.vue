<template>

  <div id="app">

    <router-view />
  </div>
</template>

<script>
import { getRouter,setRouter } from './utils/auth';

export default {
  name: "App",
  created() {
    window.addEventListener("beforeunload", () => {
      let Router =JSON.parse(getRouter())
      Router.isRouter=false;
      setRouter(JSON.stringify(Router));
    });
  },
  data() {
    return {
      title: "App",
    };
  },
};
</script>
