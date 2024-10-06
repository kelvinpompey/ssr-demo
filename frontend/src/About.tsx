import { AboutRouteProps } from "./generated";

function About({ content }: AboutRouteProps) {
  return (
    <div className="h-screen items-center flex-col flex justify-center bg-zinc-800">
      <div className=" text-3xl text-purple-400">Hello</div>
      <div className="dark:text-white">{content}</div>

      <a href="/" className="text-zinc-500 hover:underline">
        Home
      </a>
    </div>
  );
}

export default About;
