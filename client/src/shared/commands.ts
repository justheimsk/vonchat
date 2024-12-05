import {vonchat} from "@/lib/Application";

export default () => {
  const hello_world = () => {
    alert("Hello World!");
  }

  vonchat.cmdRegistry.register("hello_world", "Simple hello world command.", hello_world)
}
