import React, {useState} from 'react';
import PropTypes from 'prop-types';

import TableOfContents from '../Docs/TOC';
import Sidebar from '../Sidebar';
import Header from '../Header';
import {Children, Main, Title, WarnBanner, Wrapper} from './styles';

export default function Layout({
  children,
  disableTableOfContents,
  title,
  headings,
}) {
  const [isMenuOpen, setMenuOpen] = useState(false);
  const disableTOC =
    disableTableOfContents === true || !headings || headings.length === 0;

  function handleMenuOpen() {
    setMenuOpen(!isMenuOpen);
  }

  return (
    <>
      <WarnBanner>⚠️ You are viewing the documentation for an earlier version (v1) of golangci-lint ⚠️</WarnBanner>
      <Sidebar isMenuOpen={isMenuOpen} />
      <Header handleMenuOpen={handleMenuOpen} isMenuOpen={isMenuOpen} />
      <Wrapper isMenuOpen={isMenuOpen}>
        {title && <Title>{title}</Title>}
        <Main disableTOC={disableTOC}>
          {!disableTOC && <TableOfContents headings={headings} />}
          <Children hasTitle={title}>{children}</Children>
        </Main>
      </Wrapper>
    </>
  );
}

Layout.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]).isRequired,
  disableTableOfContents: PropTypes.bool,
  title: PropTypes.string,
  headings: PropTypes.array,
};

Layout.defaultProps = {
  disableTableOfContents: false,
  title: '',
  headings: null,
};
