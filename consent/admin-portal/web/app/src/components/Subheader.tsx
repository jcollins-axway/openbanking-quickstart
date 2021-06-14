import React, { ReactNode, CSSProperties } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    backgroundColor: "#002D4C",
    padding: "32px 0",
    boxSizing: "border-box",
    textAlign: "center",
  },
  content: {
    maxWidth: 850,
    margin: "auto",
    color: "white",
  },
  title: {
    ...theme.custom.heading2,
  },
  info: {
    marginTop: 12,
    ...theme.custom.body2,
  },
}));

type Props = {
  title: string | ReactNode;
  children?: ReactNode;
  containerStyle?: CSSProperties;
  contentStyle?: CSSProperties;
};

function Subheader({ title, children, containerStyle, contentStyle }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container} style={containerStyle}>
      <div className={classes.content} style={contentStyle}>
        <div className={classes.title}>{title}</div>
        {children && <div className={classes.info}>{children}</div>}
      </div>
    </div>
  );
}

export default Subheader;