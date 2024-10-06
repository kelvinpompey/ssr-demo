import { Button } from "@/components/ui/button";
import { AboutRouteProps } from "./generated";
import { Input } from "./components/ui/input";

function About({ content }: AboutRouteProps) {
  return (
    <div className="h-screen items-center flex-col flex justify-center">
      <div className=" text-3xl text-purple-400">Hello</div>
      <div className="dark:text-white">{content}</div>

      <div className="max-w-[300px] gap-3 flex flex-col">
        <a href="/" className="text-zinc-500 hover:underline">
          Home
        </a>

        <Button>Hello from ShadCN</Button>
        <Input placeholder="Email" />
      </div>
    </div>
  );
}

export default About;
