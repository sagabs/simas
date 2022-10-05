import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "../../../components/sidebaradmin/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamaadmin/halamanutama";

export default function home() {
  return (
    <div className={style.home}>
      <Sidebar />
      <div className={style.homeContainer}>
        <div className={style.content}>
          <HalamanUtama />
        </div>
      </div>
    </div>
  );
}
