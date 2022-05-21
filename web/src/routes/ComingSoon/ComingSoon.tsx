import React from 'react';
import classes from './ComingSoon.module.css';
import { FaGithub } from 'react-icons/fa';

const ComingSoon: React.FC = () => {
  return (
    <div className={classes.wrapper}>
      <div className={classes.container}>
        <img
          className={classes.image}
          src="/logo.svg"
          alt="DevRoot Logo"
          loading="lazy"
        />
        <h1 className={classes.title}>DEV ROOT</h1>
        <p className={classes.infoText}>Social Media Platform For Developers</p>
        <p className={classes.comingSoon}>Coding in progress...</p>
        <a
          href="https://github.com/olafszymanski/devroot"
          target="_blank"
          rel="noreferrer"
          className={classes.callToAction}
        >
          <span className={classes.icon}>
            <FaGithub />
          </span>{' '}
          Drop a Star
        </a>
      </div>
    </div>
  );
};

export default ComingSoon;
