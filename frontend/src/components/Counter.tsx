import { Button } from "@/components/ui/button";

type Props = {
  count: number;
  increment: () => void;
};
export default function Counter({ count, increment }: Props) {
  return (
    <Button onClick={increment} className="">
      Count is {count}
    </Button>
  );
}
