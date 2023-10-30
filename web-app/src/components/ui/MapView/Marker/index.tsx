import IconDelete from "../../../../icons/delete";
import "./styles.css"

interface MarkerLayoutProps {
  name?: string;
}

export default function MarkerLayout({ name }: MarkerLayoutProps) {
  return (
    <div className="container_marker">
      <div className="circle"></div>
      <IconDelete color="#007AFF" />
      <p className="name">#{name}</p>
    </div>
  );
}
