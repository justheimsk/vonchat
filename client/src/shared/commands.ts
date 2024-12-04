import {vonchat} from "@/lib/Application";

export default () => {
  const hello_world = () => {
    alert("Hello World!");
  }

  vonchat.cmdRegistry.register("hello_world", "Simple Hello World Command.", hello_world)
}
