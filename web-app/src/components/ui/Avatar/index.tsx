import "./styles.css";

interface AvatarProps {
  imgAvatar: string;
  nameAvatar: string;
}

export default function Avatar({ imgAvatar, nameAvatar }: AvatarProps) {
  return (
    <span className="container_avatar">
      <img src={imgAvatar} alt={nameAvatar} className="img_avatar" />
      <p className="text_avatar">{nameAvatar}</p>
    </span>
  );
}
