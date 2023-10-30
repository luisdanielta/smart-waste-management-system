import IconDelete from "../../../../icons/delete";
import "./styles.css"

interface MarkerLayoutProps {
  name: string;
}

export default function MarkerLayout({ name }: MarkerLayoutProps) {
  return (
    <div className="container_marker">
      <IconDelete color="#007AFF" />
        <p>#{name}</p>
    </div>
  );
}
