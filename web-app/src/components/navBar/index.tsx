import "./styles.css";
import Avatar from "../ui/Avatar";
import IconNavigation from "../../icons/navigation";
import IconSettings from "../../icons/settings";

export default function NavBar() {
  return (
    <nav className="container_navbar">
      <div className="container_navbar_elements">
        <IconNavigation color="#007AFF" className="icon_button" />
        <span className="container_avatar_navbar">
          <span className="container_avatar_navbar">
            {/* user icon view */}
            <p>Admin</p>
            <Avatar
              imgAvatar="https://avatars.githubusercontent.com/u/72309895?v=4"
              nameAvatar="luisdanielta"
            />
          </span>
          <IconSettings color="#007AFF" className="icon_button" />
        </span>
      </div>
    </nav>
  );
}
