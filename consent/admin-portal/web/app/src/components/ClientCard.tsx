import React, { useState } from "react";
import { makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import clsx from "clsx";

import { ClientType, ConsentStatus, getDate } from "./utils";
import RevokeDrawer from "./ThirdPartyProvidersView/RevokeDrawer";

const useStyles = makeStyles((theme: Theme) => ({
  card: {
    background: "#FFFFFF",
    padding: "12px 24px",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    marginBottom: 16,
  },
  cardDisabled: {
    backgroundColor: "#ECECEC",
  },
  date: {
    ...theme.custom.caption,
  },
  cardName: {
    display: "flex",
    alignItems: "center",
  },
  avatar: {
    background: "#FCFCFF",
    border: "1px solid #626576",
    color: "#626576",
    width: 48,
    height: 48,
    marginRight: 12,
    borderRadius: 4,
  },
  name: {
    ...theme.custom.heading3,
  },
  buttonRoot: {
    ...theme.custom.body2,
    textTransform: "capitalize",
    color: "#002D4C",
    backgroundColor: "#FCFCFF",
    border: "1px solid #002D4C",
    padding: "8px 24px",
  },
  buttonRootDisabled: {
    backgroundColor: "#DC1B37",
    border: "1px solid #DC1B37",
    color: "white !important",
  },
}));

function getAuthorisedDate(client) {
  const accountAccessConsent = client?.consents?.find(
    (v) => v.type === "account_access"
  );
  if (accountAccessConsent) {
    const date = accountAccessConsent?.account_access_consent?.CreationDateTime;
    return getDate(date);
  }
  return "N/A";
}

interface PropTypes {
  client: ClientType;
  onRevokeClient: (id: string) => void;
}

export default function ClientCard({ client, onRevokeClient }: PropTypes) {
  const classes = useStyles();
  const date = getAuthorisedDate(client);
  const [openDrawer, setOpenDrawer] = useState(false);
  const status = client?.mainStatus;

  return (
    <div
      className={clsx(
        classes.card,
        status === ConsentStatus.Inactive && classes.cardDisabled
      )}
    >
      <div className={classes.cardName}>
        <Avatar variant="square" className={classes.avatar}>
          {client?.client_name[0]?.toUpperCase()}
        </Avatar>
        <div>
          <div className={classes.name}>{client?.client_name}</div>
          {date !== "N/A" && (
            <div className={classes.date}>Connected {date}</div>
          )}
        </div>
      </div>
      {status !== ConsentStatus.Inactive ? (
        <Button
          variant="outlined"
          classes={{ root: classes.buttonRoot }}
          onClick={() => setOpenDrawer(true)}
        >
          Revoke
        </Button>
      ) : (
        <Button
          classes={{
            root: clsx(classes.buttonRoot, classes.buttonRootDisabled),
          }}
          disabled
        >
          Revoked
        </Button>
      )}

      {openDrawer && (
        <RevokeDrawer
          handleClose={() => setOpenDrawer(false)}
          onConfirm={() =>
            client?.client_id &&
            onRevokeClient &&
            onRevokeClient(client?.client_id)
          }
          client={client}
        />
      )}
    </div>
  );
}
