import moment from "moment";
import * as React from "react";
import {
  Button,
  Card,
  Feed,
  Header,
  Icon,
  List,
  SemanticCOLORS
} from "semantic-ui-react";
import { HTTP_URL } from "../../config";
import { File } from "../../graphql/models";
import { XClipboard } from "../form";
import XCreateLinkModal from "../link/XLinkCreateModal";
import XFileUploadModal from "./XFileUploadModal";

const XFileCard = (f: File) => {
  let colors: SemanticCOLORS[] = [
    "olive",
    "green",
    "teal",
    "blue",
    "violet",
    "purple",
    "pink"
  ];

  return (
    <Card fluid>
      <Card.Content>
        <Button.Group floated="right">
          <XCreateLinkModal file={f.id} />
          <XFileUploadModal
            fileName={f.name}
            button={{ basic: true, color: "blue", icon: "cloud upload" }}
          />{" "}
          <Button
            basic
            color="blue"
            icon="cloud download"
            href={HTTP_URL + "/cdn/download/" + f.name}
          />
        </Button.Group>
        <Card.Header>{f.name}</Card.Header>
        <Card.Meta>{f.size} bytes</Card.Meta>
        <Card.Description>
          <Header
            sub
            disabled={!f.links || f.links.length < 1}
            style={{ marginTop: "5px" }}
          >
            <Header.Content>
              {f.links && f.links.length > 0
                ? "Links (" + f.links.length + " total)"
                : "No Active Links"}
            </Header.Content>
          </Header>
          <Feed style={{ maxHeight: "25vh", overflowY: "auto" }}>
            {f.links && f.links.length > 0 ? (
              f.links.map((link, index) => (
                <Feed.Event key={index}>
                  <Feed.Label>
                    <Icon
                      fitted
                      name="linkify"
                      color={colors[Math.floor(Math.random() * colors.length)]}
                    />
                  </Feed.Label>
                  <Feed.Content>
                    <Feed.Summary>
                      <List.Header>
                        <XClipboard value={"/l/" + link.alias}>
                          {"/l/" + link.alias}
                        </XClipboard>
                      </List.Header>
                      <Feed.Date>
                        {link.expirationTime
                          ? "Expires " + moment().to(link.expirationTime)
                          : "Never expires"}
                      </Feed.Date>
                    </Feed.Summary>
                    <Feed.Meta>
                      {link.clicks && link.clicks > 0
                        ? link.clicks + " Clicks left"
                        : "Unlimited clicks"}
                    </Feed.Meta>
                  </Feed.Content>
                </Feed.Event>
              ))
            ) : (
              <span />
            )}
          </Feed>
        </Card.Description>
      </Card.Content>
      <Card.Content extra>
        Created: {moment(f.creationTime).fromNow()}
        <br />
        Last Modified: {moment(f.lastModifiedTime).fromNow()}
      </Card.Content>
    </Card>
  );
};

export default XFileCard;
