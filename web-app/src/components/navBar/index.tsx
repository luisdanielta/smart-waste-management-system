import "./styles.css";
import Avatar from "../ui/Avatar";

export default function NavBar() {
  return (
    <nav className="container_navbar">
      <span
        style={{
          color: "#007AFF",
        }}
      >
        <p>Menu</p>
      </span>

      <span
        style={{
          color: "#007AFF",
        }}
      >
        <p>Home</p>
      </span>

      <span className="container_avatar_navbar">
        {/* user icon view */}
        <p>Admin</p>
        <Avatar
          imgAvatar="https://avatars.githubusercontent.com/u/72309895?v=4"
          nameAvatar="luisdanielta"
        />
        <span
          style={{
            color: "#007AFF",
          }}
        >
          <p>Settings</p>
        </span>
      </span>
    </nav>
  );
}
